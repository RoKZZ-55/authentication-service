package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Error fills the response body with an error and redirects the data to the responder
func Error(w http.ResponseWriter, code int, err error) {
	if errors.Is(err, errors.New("internal server error")) {
		code = http.StatusInternalServerError
	}
	Respond(w, code, map[string]string{"error": err.Error()})
}

// Respond returns a response to a client request in the form of status code and response body
func Respond(w http.ResponseWriter, code int, data interface{}) {
	// return code status to client
	w.WriteHeader(code)

	// body return to client if not nil
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
