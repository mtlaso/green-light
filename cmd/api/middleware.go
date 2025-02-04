package main

import (
	"fmt"
	"net/http"
)

// recoverPanic recovers from a panic, logs the details, and sends a 500 Internal error.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event of a panic
		// as Go unwinds the stack).
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)

}
