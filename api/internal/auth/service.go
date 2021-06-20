package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fudge/snooker/internal/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type Service interface {
	Register(context.Context, *entity.User) error
	Authenticate(context.Context, string, string) (*jwt.Token, error)
}

type AuthService struct {
	Users entity.UserRepository
}

// Create a new Authentication Service
func New(u entity.UserRepository) Service {
	return &AuthService{Users: u}
}

// Register
// Register a new User with validation and then store the User in the
// storage layer of the application
// @TODO Send email to the User regarding their new account
func (a AuthService) Register(ctx context.Context, u *entity.User) error {
	if err := validateNewUser(*u, ctx, a.Users); err != nil {
		return ErrValidationFailed{
			Errors: err.(validation.Errors),
		}
	}

	if err := a.Users.AddUser(ctx, *u); err != nil {
		return errors.New("could not store user")
	}

	return nil
}

// Authenticate a User against a set of credentials
// When a user is not found an error will be returned
func (a AuthService) Authenticate(ctx context.Context, email string, password string) (*jwt.Token, error) {
	u, err := a.Users.FindUserByCredentials(ctx, email, password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", u.Id),
		ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
	})

	return token, nil
}
