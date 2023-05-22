package main

import (
	"flag"
	"os"

	"github.com/forstes/besafe-go/customer/pkg/hash"
	"github.com/forstes/besafe-go/customer/pkg/store/postgres"
	http "github.com/forstes/besafe-go/customer/services/customer/internal/delivery/http"
	v1 "github.com/forstes/besafe-go/customer/services/customer/internal/delivery/http/v1"
	"github.com/forstes/besafe-go/customer/services/customer/internal/repository"
	"github.com/forstes/besafe-go/customer/services/customer/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	dbConnCfg := postgres.ConnectionConfig{}
	httpServerCfg := http.ServerConfig{}

	err := godotenv.Load(".env")
	if err != nil {
		println("Failed to load .env file", err)
		os.Exit(1)
	}

	flag.IntVar(&httpServerCfg.Port, "http-port", 8000, "HTTP server port")
	flag.StringVar(&httpServerCfg.ReadTimeout, "http-read-timeout", "10s", "HTTP read timeout")
	flag.StringVar(&httpServerCfg.WriteTimeout, "http-write-timeout", "30s", "HTTP write timeout")
	flag.StringVar(&httpServerCfg.IdleTimeout, "http-idle-timeout", "1m", "HTTP idle timeout")

	flag.IntVar(&dbConnCfg.Port, "pg-port", 5432, "Postgres port")
	flag.StringVar(&dbConnCfg.Host, "pg-host", "localhost", "Postgres host")
	flag.StringVar(&dbConnCfg.User, "pg-user", os.Getenv("PG_USER"), "Postgres user")
	flag.StringVar(&dbConnCfg.Password, "pg-password", os.Getenv("PG_PASSWORD"), "Postgres password")
	flag.StringVar(&dbConnCfg.DbName, "pg-db-name", os.Getenv("PG_DB_NAME"), "Postgres DB name")
	flag.IntVar(&dbConnCfg.MaxOpenConnections, "pg-max-open-conns", 15, "Postgres max open connections")
	flag.StringVar(&dbConnCfg.MaxIdleTime, "pg-max-idle-time", "15m", "Postgres max connection idle time")
	flag.Parse()

	db, err := postgres.OpenDB(dbConnCfg)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	defer db.Close()

	// TODO Add Logger
	println("Connected to Postgres DB")

	passwordHasher := hash.NewSHA256Hasher("Bruhable")
	userRepository := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepository, passwordHasher)

	httpServerV1 := http.NewHttpServer(v1.NewRouter(userService).GetRoutes(), httpServerCfg)
	err = httpServerV1.Serve()
	if err != nil {
		println("Failed to start HTTP server", err)
	}
}
