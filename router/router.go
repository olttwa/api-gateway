package router

import (
	"rgate/docker"
	"rgate/handler"
	"rgate/middleware"
	"rgate/model"

	"github.com/gorilla/mux"
)

type config interface {
	Routes() []model.Route
	DefaultBody() []byte
	DefaultCode() int
}

func New(cfg config) *mux.Router {
	r := mux.NewRouter()

	d := docker.New()
	for _, route := range cfg.Routes() {
		h := handler.Backend(d, route.MatchLabels)
		r.PathPrefix(route.PathPrefix).Handler(h)
	}

	stats := middleware.StatsMw()
	r.Use(stats.Middleware)
	r.Handle("/stats", handler.Stats(stats))

	// Refrain from using Mux Router's NotFoundHandler for default
	// responses because stats middleware doesn't intercept them.
	r.PathPrefix("").Handler(handler.Default(cfg.DefaultBody(), cfg.DefaultCode()))

	return r
}
