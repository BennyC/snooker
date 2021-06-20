package auth

import validation "github.com/go-ozzo/ozzo-validation/v4"

type ErrValidationFailed struct {
	Errors validation.Errors
}

func (v ErrValidationFailed) Error() string {
	return "failed validation"
}
