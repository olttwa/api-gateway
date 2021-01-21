package handler

import (
	"log"
	"net/http"
	"rgate/config"
)

func Default() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(config.DefaultResponseStatusCode())
		if _, err := w.Write(config.DefaultResponseBody()); err != nil {
			log.Printf("error writing default response: %s\n", err)
		}
	})
}
