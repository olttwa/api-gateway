package main

import (
	"rgate/config"
	"rgate/gateway"
	"rgate/router"
)

func main() {
	cfg := config.Load()

	r := router.New(cfg)
	gateway.Serve(cfg.Port(), r)
}
