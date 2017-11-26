package helpers

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Message string `json:"message"`
}

// Error is helper for err response
func Error(w http.ResponseWriter, msg string, code int) {
	err := jsonError{
		Message: msg,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	errJSON, _ := json.Marshal(&err)
	w.Write(errJSON)
	// fmt.Fprintln(w, error)
}
