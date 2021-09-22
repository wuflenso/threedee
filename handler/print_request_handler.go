package handler

import (
	"errors"
	"fmt"
	"net/http"
	"threedee/utility/response"

	"github.com/julienschmidt/httprouter"
)

/*
 * THIRD LAYER => All .go files in handler directory
 * This file is the third layer of the service. It's objective is to handle specific
 * business usecases that might contain the business logic
 */

type RequestHandler struct{}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

// handle GET /requests
func (h *RequestHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	fmt.Fprintf(w, "hallo")
	data := "kosong"
	return response.WriteSuccess(w, data, "success")

}

// handle GET /requests/:id
func (h *RequestHandler) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	return response.WriteBadRequest(w, errors.New("kaco inputnya"))
}
