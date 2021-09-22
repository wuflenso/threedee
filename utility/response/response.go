package response

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	HttpStatus int         `json:"http_status"`
}

func Respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, data interface{}, message string) error {
	meta := Meta{
		Message:    message,
		Data:       data,
		HttpStatus: http.StatusOK,
	}

	Respond(w, meta, http.StatusOK)
	return nil
}

func WriteError(w http.ResponseWriter, err error) error {
	meta := Meta{
		Message:    err.Error(),
		HttpStatus: http.StatusInternalServerError,
	}

	Respond(w, meta, http.StatusInternalServerError)
	return err
}

func WriteBadRequest(w http.ResponseWriter, err error) error {
	meta := Meta{
		Message:    "Bad Request",
		HttpStatus: http.StatusBadRequest,
	}

	Respond(w, meta, http.StatusOK)
	return err
}
