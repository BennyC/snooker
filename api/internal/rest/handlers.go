package rest

import (
	h "net/http"

	"github.com/fudge/snooker/internal/auth"
	"github.com/fudge/snooker/internal/entity"
	"github.com/spf13/viper"
)

// Register HTTP handler
func Register(a auth.Authentication) h.HandlerFunc {
	return func(w h.ResponseWriter, r *h.Request) {
		var u entity.User
		if err := ReadRequest(r, &u); err != nil {
			InvalidRequestResponse(w, err)
			return
		}

		// Change the HTTP response based on the error returned
		// by the Register function
		if err := a.Register(r.Context(), &u); err != nil {
			switch e := err.(type) {
			case auth.ErrValidationFailed:
				UnprocessableEntityResponse(w, e.Errors)
				return
			default:
				InternalServerErrorResponse(w, e)
				return
			}
		}

		w.WriteHeader(h.StatusCreated)
	}
}

// Authenticate HTTP handler
func Authenticate(a auth.Authentication) h.HandlerFunc {
	type AuthRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w h.ResponseWriter, r *h.Request) {
		var req AuthRequest
		if err := ReadRequest(r, &req); err != nil {
			InvalidRequestResponse(w, err)
			return
		}

		token, err := a.Authenticate(r.Context(), req.Email, req.Password)
		if err != nil {
			UnauthorizedResponse(w, nil)
			return
		}

		// @TODO
		// Need to return a refresh token, as well as an access token
		key := viper.GetString("JWT_SECRET_KEY")
		json, err := token.SignedString([]byte(key))
		if err != nil {
			InternalServerErrorResponse(w, err)
			return
		}

		w.WriteHeader(h.StatusOK)
		w.Write([]byte(json))
	}
}
