package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes defines the application routes.
func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)

	// Wrap the router with the panic recovery middleware.
	// Will only recover panics that happen in the same goroutine
	// that executed the `recoverPanic` middleware.
	return app.recoverPanic(router)
}
