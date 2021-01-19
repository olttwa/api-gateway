package main

import (
	"rgate/config"
	"rgate/docker"
	"rgate/gateway"
	"rgate/router"
	"rgate/utils"
)

func main() {
	utils.SeedRandom()
	config.Load()
	d := docker.Initialize()

	r := router.New(d)
	gateway.Serve(r)
}
