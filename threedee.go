package threedee

import (
	"net/http"
	"threedee/handler"
	m "threedee/middleware"
	"threedee/repository"
	"threedee/utility/normalizer"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

/*
 * SECOND LAYER => Threedee.go or (whateverurappnameis).go
 *
 * This file is the second layer of the service. It's objective is to route endpoints to
 * corresponding handlers.
 *
 * This file contains instantiation of the app, which
 * includes the declaration and instantiation of the services' parts like router, handler, repo,
 * service, and supporting packages like CORS and httprouter.
 *
 * Endpoints usually route to handler package which contains all the logic to handle
 * specific usecases. So, limit this threedee.go into routing and related configs like
 * auth and database setup.
 */

type Threedee struct {
	Router http.Handler
}

func NewThreedee() *Threedee {

	corsConfig := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
	})

	router := httprouter.New()

	// We input the repo here, not the interface. The interface is for contraint purpose only
	rep := repository.NewPrintRequestRepository()
	norm := normalizer.NewPrintRequestNormalizer()
	rh := handler.NewRequestHandler(rep, norm)
	router.GET("/print-requests", m.Middleware(rh.Index))
	router.GET("/print-requests/:id", m.Middleware(rh.Show))
	router.POST("/print-requests", m.Middleware(rh.Create))
	router.PUT("/print-requests/:id", m.Middleware(rh.Update))
	router.PUT("/print-requests/:id/status", m.Middleware(rh.ChangeStatus))
	router.DELETE("/print-requests/:id", m.Middleware(rh.Delete))

	return &Threedee{corsConfig.Handler(router)}
}
