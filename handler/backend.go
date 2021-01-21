package handler

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"rgate/docker"
	"rgate/utils"
)

type backend struct {
	containers []docker.Container
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

func Backend(d docker.Client, ml []string) http.Handler {
	ctx := context.Background()
	c, err := d.ListContainers(ctx, ml)
	if err != nil {
		log.Fatalf("error listing containers: %s", err)
	}
	return backend{c}
}
