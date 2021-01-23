package utils

import (
	"log"
	"net/http"
)

const (
	BackendErrorCode = http.StatusServiceUnavailable
)

var ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf(r.URL.Path+" error: %s\n", err)
	w.WriteHeader(BackendErrorCode)
}
