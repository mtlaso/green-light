package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// envelope is used to hold a JSON object inside some value.
//
// Example:
//
//		{
//		 "envelope-key": {
//	      ...JSON object...
//		}
//
// }
type envelope map[string]any

// readIDParam reads the id parameter from the URL and converts it to an integer.
// If the id parameter cannot be parsed to an integer, or is less than 1, it returns an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// writeJSON responds to the request with a JSON payload.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newline to the JSON response body, for better readability.
	js = append(js, '\n')

	// Set headers.
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
