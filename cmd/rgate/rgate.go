package main

import (
	"rgate/config"
	"rgate/docker"
	"rgate/router"
)

func main() {
	config.Load()
	d := docker.Initialize()
	router.New(d)
}
