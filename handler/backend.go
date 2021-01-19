package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"rgate/docker"
	"rgate/utils"
)

func Backend(c []docker.Container) http.Handler {
	host := randomBackend(c)
	url := url.URL{
		Scheme: "http",
		Host:   host,
	}
	rp := httputil.NewSingleHostReverseProxy(&url)
	return rp
}

func randomBackend(c []docker.Container) string {
	if len(c) == 0 {
		return ""
	}

	i := utils.RandInt(len(c))
	return c[i].IP + ":" + c[i].Port
}
