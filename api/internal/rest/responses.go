package rest

import (
	"bytes"
	"encoding/json"
	h "net/http"
)

func UnauthorizedResponse(w h.ResponseWriter, err error) {
	JSON(w, h.StatusUnauthorized, nil)
}

func InternalServerErrorResponse(w h.ResponseWriter, err error) {
	JSON(w, h.StatusInternalServerError, nil)
}

func InvalidRequestResponse(w h.ResponseWriter, err error) {
	JSON(w, h.StatusBadRequest, err)
}

func UnprocessableEntityResponse(w h.ResponseWriter, err error) {
	type response struct {
		Errors error `json:"errors"`
	}

	JSON(w, h.StatusUnprocessableEntity, response{Errors: err})
}

func JSON(w h.ResponseWriter, code int, v interface{}) {
	b := new(bytes.Buffer)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(b).Encode(v); err != nil {
		w.WriteHeader(h.StatusInternalServerError)
	}

	w.WriteHeader(code)
	_, _ = w.Write(b.Bytes())
}
