package repository

import (
	"context"

	"github.com/forstes/besafe-go/customer/services/customer/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetOne(ctx context.Context, id string) *domain.User
	GetMany(ctx context.Context, limit, offset int) []*domain.User
}

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{db: db}
}

func (s *userRepo) Create(ctx context.Context, user domain.User) error {
	return nil
}

func (s *userRepo) GetOne(ctx context.Context, id string) *domain.User {
	return nil
}

func (s *userRepo) GetMany(ctx context.Context, limit, offset int) []*domain.User {
	return nil
}
