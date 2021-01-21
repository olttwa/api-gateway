package router

import (
	"rgate/config"
	"rgate/docker"
	"rgate/handler"
	"rgate/middleware"

	"github.com/gorilla/mux"
)

func New(d docker.Client) *mux.Router {
	r := mux.NewRouter()

	for _, route := range config.Routes() {
		h := handler.Backend(d, route.MatchLabels)
		r.PathPrefix(route.PathPrefix).Handler(h)
	}

	stats := middleware.StatsMw()
	r.Use(stats.Middleware)
	r.Handle("/stats", handler.StatsHandler(stats))

	// Refrain from using Mux Router's NotFoundHandler for default
	// responses because stats middleware doesn't intercept them.
	r.PathPrefix("").Handler(handler.Default())

	return r
}
