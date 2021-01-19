package router

import (
	"context"
	"log"
	"net/http/httputil"
	"net/url"
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

		url := url.URL{
			Scheme: "http",
			Host:   c[0].IP + ":" + c[0].Port,
		}
		rp := httputil.NewSingleHostReverseProxy(&url)
		r.PathPrefix(route.PathPrefix).Handler(rp)
	}
	return r
}
