package handler

import (
	"log"
	"net/http"
)

func Default(body []byte, code int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		if _, err := w.Write(body); err != nil {
			log.Printf("error writing default response: %s\n", err)
		}
	})
}
