package handler

import (
	"errors"
	"net/http"
	"strconv"
	"threedee/repository"
	"threedee/utility/response"

	"github.com/julienschmidt/httprouter"
)

/*
 * THIRD LAYER => All .go files in handler directory
 * This file is the third layer of the service. It's objective is to handle specific
 * business usecases that might contain the business logic
 */

type RequestHandler struct {
	Repo *repository.PrintRequestRepository
}

func NewRequestHandler(repo *repository.PrintRequestRepository) *RequestHandler {
	return &RequestHandler{repo}
}

// handle GET /requests
func (h *RequestHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	data, err := h.Repo.GetAll()
	if err != nil {
		return response.WriteInternalServerError(w, err)
	}
	return response.WriteSuccess(w, data, "success")

}

// handle GET /requests/:id
func (h *RequestHandler) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return response.WriteBadRequestError(w, errors.New("id is not a number"))
	}

	data, err := h.Repo.GetById(id)
	if err != nil {
		return response.WriteInternalServerError(w, err)
	}
	if data.Id == 0 {
		return response.WriteNotFoundError(w, errors.New("record not found"))
	}

	return response.WriteSuccess(w, data, "success")
}
