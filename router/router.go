package router

import (
	"context"
	"log"
	"rgate/config"
	"rgate/docker"
	"rgate/handler"

	"github.com/gorilla/mux"
)

func New(d docker.Client) *mux.Router {
	ctx := context.Background()
	r := mux.NewRouter()
	r.NotFoundHandler = handler.NotFound()

	for _, route := range config.Routes() {
		c, err := d.ListContainers(ctx, route.MatchLabels)
		if err != nil {
			log.Fatalf("error listing containers: %s", err)
		}

		h := handler.Backend(c)
		r.PathPrefix(route.PathPrefix).Handler(h)
	}
	return r
}
