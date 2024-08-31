package http_utils

import (
	"encoding/json"
	"net/http"
)

func GetHost(r *http.Request) string {
	if r.TLS == nil {
		return "http://" + r.Host
	}
	return "https://" + r.Host
}

// RespondWithError returns an error to the body with the error message along with the status code
func RespondWithError(w http.ResponseWriter, code int, msg interface{}) {
	RespondWithJSON(w, code, msg)
}

// RespondWithJSON return the data to the body along with the status code
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
