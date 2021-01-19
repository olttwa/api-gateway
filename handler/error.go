package handler

import (
	"log"
	"net/http"
)

var ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf(r.URL.Path+" error: %s\n", err)
	w.WriteHeader(http.StatusServiceUnavailable)
}
