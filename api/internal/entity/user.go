package entity

import (
	"context"
	"time"
)

type User struct {
	Id        int
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
}

type UserRepository interface {
	AddUser(context.Context, User) error
	FindUserByEmail(context.Context, string) (*User, error)
	FindUserByCredentials(context.Context, string, string) (*User, error)
}
