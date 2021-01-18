package config

import (
	"flag"
	"log"
)

type config struct {
	port   int
	routes []Route
}

type Route struct {
	PathPrefix  string
	MatchLabels map[string]string
}

var cfg config

func Port() int {
	return cfg.port
}

func Routes() []Route {
	return cfg.routes
}

func Load() {
	cfgFile := flag.String("config", "config.yml", "config filename")
	port := flag.Int("port", 8080, "port for receiving traffic")

	flag.Parse()

	routes, err := ParseRoutes(*cfgFile)
	if err != nil {
		log.Fatalf("Config loading failed: %s", err.Error())
	}

	cfg = config{
		port:   *port,
		routes: routes,
	}
}
