package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fudge/snooker/internal/crypto"
	"github.com/fudge/snooker/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Db *sql.DB
}

var (
	ErrNoUserFound = errors.New("no user found")
)

func (u *Users) AddUser(c context.Context, user entity.User) error {
	stmt, err := u.Db.Prepare("INSERT INTO users (email, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	p, err := crypto.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(user.Email, p); err != nil {
		return err
	}

	return nil
}

func (u *Users) FindUserByEmail(c context.Context, e string) (*entity.User, error) {
	var user entity.User
	row := u.Db.QueryRow("SELECT id, email, password, created_at FROM users WHERE email = $1", e)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &user, err
	}
}

func (u *Users) FindUserByCredentials(c context.Context, email string, password string) (*entity.User, error) {
	user, err := u.FindUserByEmail(c, email)
	if err != nil || user == nil {
		return nil, ErrNoUserFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrNoUserFound
	}

	return user, nil
}
