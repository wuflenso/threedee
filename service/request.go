package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
 * THIRD LAYER => All .go files in service directory
 * This file is the third layer of the service. It's objective is to handle specific
 * business usecases that might contain the logic/remote api call/
 * db data retrieval/modification
 */

type RequestHandler struct{}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

// handle GET /requests
func (h *RequestHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hallo")
}

// handle GET /requests/:id
func (h *RequestHandler) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hallo lagi %s", p)
}
