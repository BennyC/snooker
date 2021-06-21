package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fudge/snooker/internal/rest"
	"github.com/fudge/snooker/internal/storage/postgres"
	"github.com/fudge/snooker/internal/testutils"
)

func TestIntegrationAuthenticationBadRequest(t *testing.T) {
	db := testutils.NewDatabase()
	down := testutils.Migrate(db)
	defer down()

	s := rest.New(&postgres.Users{Db: db})
	req := httptest.NewRequest("POST", "http://example.com/register", nil)
	res := httptest.NewRecorder()

	s.Routes().ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Errorf("expected status: %d. got: %d", http.StatusBadRequest, res.Code)
	}

}

func TestIntegrationAuthenticationValidation(t *testing.T) {
	db := testutils.NewDatabase()
	down := testutils.Migrate(db)
	defer down()

	testCases := []struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		{
			Email:    "example@example.com",
			Password: "",
		},
		{
			Email:    "",
			Password: "test1234!",
		},
		{
			Email:    "",
			Password: "",
		},
	}

	s := rest.New(&postgres.Users{Db: db})
	for _, tc := range testCases {
		jsonStr, _ := json.Marshal(tc)
		req := httptest.NewRequest("POST", "http://example.com/register", bytes.NewBuffer(jsonStr))
		res := httptest.NewRecorder()

		s.Routes().ServeHTTP(res, req)

		if http.StatusUnprocessableEntity != res.Code {
			t.Errorf("expected status: %d. got: %d", http.StatusUnprocessableEntity, res.Code)
		}
	}
}

func TestIntegrationAuthenticationSuccess(t *testing.T) {
	db := testutils.NewDatabase()
	down := testutils.Migrate(db)
	defer down()

	testCases := []struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		{
			Email:    "example@example.com",
			Password: "test1234!",
		},
		{
			Email:    "ben@example.com",
			Password: "test1234!",
		},
	}

	s := rest.New(&postgres.Users{Db: db})
	for _, tc := range testCases {
		jsonStr, _ := json.Marshal(tc)
		req := httptest.NewRequest("POST", "http://example.com/register", bytes.NewBuffer(jsonStr))
		res := httptest.NewRecorder()

		s.Routes().ServeHTTP(res, req)

		if http.StatusCreated != res.Code {
			t.Errorf("expected status: %d. got: %d", http.StatusUnprocessableEntity, res.Code)
		}
	}
}
