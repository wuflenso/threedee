package handler

import (
	"errors"
	"net/http"
	"strconv"
	print_request "threedee/interfaces/print-request"
	"threedee/utility/normalizer"
	"threedee/utility/response"

	"github.com/julienschmidt/httprouter"
)

/*
 * THIRD LAYER => Handler package
 *
 * This file is the third layer of the service. It's objective is to handle specific
 * business usecases that might contain the business logic.
 *
 * Things we usually do in this layer is something like input normalizing,
 * repo/service call, and input/output validations by calling helper packages.
 *
 * A handler utilizes either:
 * a Repository (making requests to database), or
 * a Service (making requests to other services i.e. using HTTP REST)
 */

type RequestHandler struct {
	Repo print_request.PrintRequestRepositoryInterface
	Norm *normalizer.PrintRequestNormalizer
}

func NewRequestHandler(repo print_request.PrintRequestRepositoryInterface, norm *normalizer.PrintRequestNormalizer) *RequestHandler {
	return &RequestHandler{repo, norm}
}

// handle GET /print-requests
func (h *RequestHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (int, error) {

	ctx := r.Context()
	select {
	case <-ctx.Done():
		return http.StatusRequestTimeout, ctx.Err()
	default:
	}

	data, err := h.Repo.GetAll()
	if len(data) == 0 {
		return http.StatusNotFound, response.WriteNotFoundError(w, errors.New("no records found"))
	}
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}
	return http.StatusOK, response.WriteSuccess(w, data, "success")

}

// handle GET /print-requests/:id
func (h *RequestHandler) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) (int, error) {

	ctx := r.Context()
	select {
	case <-ctx.Done():
		return http.StatusRequestTimeout, ctx.Err()
	default:
	}

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, errors.New("id is not a number"))
	}

	data, err := h.Repo.GetById(id)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}
	if data == nil || data.Id == 0 {
		return http.StatusNotFound, response.WriteNotFoundError(w, errors.New("record not found"))
	}

	return http.StatusOK, response.WriteSuccess(w, data, "success")
}

// handle POST /print-requests
func (h *RequestHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) (int, error) {

	ctx := r.Context()
	select {
	case <-ctx.Done():
		return http.StatusRequestTimeout, ctx.Err()
	default:
	}

	model, err := h.Norm.ReadAndNormalize(w, r)
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, err)
	}

	id, err := h.Repo.Insert(model)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}

	model, err = h.Repo.GetById(id)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}

	return http.StatusOK, response.WriteSuccess(w, model, "success")
}

// handle PUT /print-requests/:id
func (h *RequestHandler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) (int, error) {

	ctx := r.Context()
	select {
	case <-ctx.Done():
		return http.StatusRequestTimeout, ctx.Err()
	default:
	}

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, errors.New("id is not a number"))
	}

	model, err := h.Norm.ReadAndNormalize(w, r)
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, err)
	}

	data, err := h.Repo.GetById(id)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}
	if data.Id == 0 {
		return http.StatusNotFound, response.WriteNotFoundError(w, errors.New("record not found"))
	}
	if data.Status != "received" {
		return http.StatusBadRequest, response.WriteBadRequestError(w, errors.New("you can not edit a request that is already processed"))
	}

	model.Id = id
	_, err = h.Repo.Update(model)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}

	return http.StatusOK, response.WriteSuccess(w, model, "success")
}

// handle DELETE /print-requests/:id
func (h *RequestHandler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) (int, error) {

	ctx := r.Context()
	select {
	case <-ctx.Done():
		return http.StatusRequestTimeout, ctx.Err()
	default:
	}

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, errors.New("id is not a number"))
	}

	_, err = h.Repo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}

	return http.StatusOK, response.WriteSuccess(w, nil, "success")
}

// handle PUT /print-requests/:id/status
func (h *RequestHandler) ChangeStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) (int, error) {

	ctx := r.Context()
	select {
	case <-ctx.Done():
		return http.StatusRequestTimeout, ctx.Err()
	default:
	}

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, errors.New("id is not a number"))
	}

	model, err := h.Norm.ReadAndNormalize(w, r)
	if err != nil {
		return http.StatusBadRequest, response.WriteBadRequestError(w, err)
	}

	data, err := h.Repo.GetById(id)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}
	if data.Id == 0 {
		return http.StatusNotFound, response.WriteNotFoundError(w, errors.New("record not found"))
	}

	data.Status = model.Status
	_, err = h.Repo.Update(data)
	if err != nil {
		return http.StatusInternalServerError, response.WriteInternalServerError(w, err)
	}

	return http.StatusOK, response.WriteSuccess(w, model, "success")
}
