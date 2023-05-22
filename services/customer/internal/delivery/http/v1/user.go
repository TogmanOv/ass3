package v1

import (
	"encoding/json"
	"net/http"

	"github.com/forstes/besafe-go/customer/services/customer/internal/service"
)

type UserHandler struct {
	userService service.Users
}

func NewUserHandler(userService service.Users) *UserHandler {
	return &UserHandler{userService: userService}
}

type registerUserDTO struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Phone     string
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var dto registerUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		// TODO Handle err
		return
	}

	// TODO Validate
	input := service.UserSignUpInput{
		Email:     dto.Email,
		Password:  dto.Password,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Phone:     dto.LastName,
	}

	err := h.userService.SignUp(r.Context(), input)
	if err != nil {
		// TODO Handle err
		return
	}
	w.WriteHeader(http.StatusOK)
}
