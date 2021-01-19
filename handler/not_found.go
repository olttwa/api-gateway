package handler

import (
	"log"
	"net/http"
	"rgate/config"
)

type notFound struct {
	body       []byte
	statusCode int
}

func (f notFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(f.statusCode)
	if _, err := w.Write(f.body); err != nil {
		log.Printf("error writing default response: %s\n", err)
	}
}

func NotFound() http.Handler {
	return notFound{
		body:       config.DefaultResponseBody(),
		statusCode: config.DefaultResponseStatusCode(),
	}
}
