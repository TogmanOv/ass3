package service

import (
	"context"

	"github.com/forstes/besafe-go/customer/pkg/hash"
	"github.com/forstes/besafe-go/customer/services/customer/internal/domain"
	"github.com/forstes/besafe-go/customer/services/customer/internal/repository"
)

type UserSignUpInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Phone     string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Token struct {
	PlainText string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Token, error)
}

type userService struct {
	repo   repository.Users
	hasher hash.PasswordHasher
}

func NewUserService(repo repository.Users, hasher hash.PasswordHasher) *userService {
	return &userService{repo: repo, hasher: hasher}
}

func (s *userService) SignUp(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	// TODO Validate input
	user := domain.User{
		Email:        input.Email,
		PasswordHash: passwordHash,
		Details: domain.UserDetails{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Phone:     input.Phone,
		},
	}

	return s.repo.Create(ctx, user)
}

func (s *userService) SignIn(ctx context.Context, input UserSignInInput) (Token, error) {
	return Token{}, nil
}
