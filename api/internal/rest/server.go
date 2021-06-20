package rest

import (
	"encoding/json"
	h "net/http"

	"github.com/fudge/snooker/internal/auth"
	"github.com/fudge/snooker/internal/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	storage *storage.Storage
	auth    auth.Authentication
}

func NewServer(s *storage.Storage) *Server {
	return &Server{
		auth:    auth.AuthService{Users: s.Users},
		storage: s,
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
