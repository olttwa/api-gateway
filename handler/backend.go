package handler

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"rgate/model"
	"rgate/utils"
)

type backend struct {
	containers []model.Container
}

func (b backend) host() string {
	c := b.containers
	if len(c) == 0 {
		return ""
	}

	i := utils.RandInt(len(c))
	return c[i].IP + ":" + c[i].Port
}

func (b backend) reverseProxy() http.Handler {
	url := url.URL{
		Scheme: "http",
		Host:   b.host(),
	}
	rp := httputil.NewSingleHostReverseProxy(&url)
	rp.ErrorHandler = utils.ErrorHandler
	return rp
}

func (b backend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := b.reverseProxy()
	rp.ServeHTTP(w, r)
}

type Docker interface {
	ListContainers(context.Context, []string) ([]model.Container, error)
}

func Backend(d Docker, ml []string) http.Handler {
	ctx := context.Background()
	c, err := d.ListContainers(ctx, ml)
	if err != nil {
		log.Fatalf("error listing containers: %s", err)
	}
	return backend{c}
}
