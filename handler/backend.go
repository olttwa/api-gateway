package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"rgate/docker"
	"rgate/utils"
)

type backend struct {
	containers []docker.Container
}

func (b backend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp := b.reverseProxy()
	rp.ServeHTTP(w, r)
}

func Backend(c []docker.Container) http.Handler {
	return backend{c}
}

func (b backend) reverseProxy() http.Handler {
	url := url.URL{
		Scheme: "http",
		Host:   b.host(),
	}
	rp := httputil.NewSingleHostReverseProxy(&url)
	rp.ErrorHandler = ErrorHandler
	return rp
}

func (b backend) host() string {
	c := b.containers
	if len(c) == 0 {
		return ""
	}

	i := utils.RandInt(len(c))
	return c[i].IP + ":" + c[i].Port
}
