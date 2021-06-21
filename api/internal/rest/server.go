package rest

import (
	"encoding/json"
	h "net/http"

	"github.com/fudge/snooker/internal/auth"
	"github.com/fudge/snooker/internal/entity"
	"github.com/gorilla/mux"
)

type Server struct {
	auth auth.Service
}

func New(users entity.UserRepository) *Server {
	return &Server{
		auth: auth.AuthService{Users: users},
	}
}

func (s *Server) Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", Register(s.auth)).Methods("POST")
	r.HandleFunc("/authenticate", Authenticate(s.auth)).Methods("POST")
	return r
}

func ReadRequest(r *h.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
