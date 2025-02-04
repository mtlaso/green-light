package main

import "net/http"

// logError is a helper that logs the error message and the associated request method and URL.
func (app *application) logError(req *http.Request, err error) {
	app.logger.Error(err.Error(), "method", req.Method, "uri", req.URL.RequestURI())
}

// errorResponse is a generic helper that sends a JSON error response.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse sends a 500 Internal Server Error response.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	msg := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, msg)
}

// notFoundResponse sends a 404 Not Found response.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	msg := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, msg)
}

// badRequestResponse sends a 400 Bad Request response.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, message any) {
	app.errorResponse(w, r, http.StatusBadRequest, message)
}

// methodNotAllowed sends a 405 Method Not Allowed response.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	msg := "the requested HTTP method is not supported by this resource"
	app.errorResponse(w, r, http.StatusMethodNotAllowed, msg)
}
