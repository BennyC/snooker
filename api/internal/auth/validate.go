package auth

import (
	"context"

	"github.com/fudge/snooker/internal/entity"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

func validateNewUser(u entity.User, ctx context.Context, users entity.UserRepository) error {
	return validation.ValidateStructWithContext(
		ctx,
		&u,
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email, validation.By(uniqueEmail(users, ctx))),
	)
}

func uniqueEmail(users entity.UserRepository, ctx context.Context) validation.RuleFunc {
	return func(value interface{}) error {
		e, _ := value.(string)
		u, _ := users.FindUserByEmail(ctx, e)

		if u != nil {
			return errors.New("must be unique")
		}

		return nil
	}
}
