package threedee

import (
	"net/http"
	"threedee/handler"
	m "threedee/middleware"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

/*
 * SECOND LAYER => Threedee.go
 * This file is the second layer of the service. It's objective is to declare the Handler
 * item which in this case, "Threedee". This file contains instantiation of CORS and using
 * httprouter package to route endpoints easily.
 *
 * endpoints usually route to files inside service directory which contains all the logic to handle
 * specific usecase. So, limit this file into routing and related configs like auth and database config.
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

	rh := handler.NewRequestHandler()
	router.GET("/requests", m.Middleware(rh.Index))
	router.GET("/requests/:id", m.Middleware(rh.Show))

	return &Threedee{corsConfig.Handler(router)}
}
