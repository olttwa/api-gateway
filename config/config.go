package config

import (
	"flag"
	"log"
)

type config struct {
	port       string
	routes     []Route
	body       []byte
	statusCode int
}

type Route struct {
	PathPrefix  string
	MatchLabels []string
}

var cfg config

func Port() string {
	return cfg.port
}

func Routes() []Route {
	return cfg.routes
}

func DefaultResponseBody() []byte {
	return cfg.body
}

func DefaultResponseStatusCode() int {
	return cfg.statusCode
}

func Load() {
	cfgFile := flag.String("config", "config.yml", "config filename")
	port := flag.String("port", "8080", "port for receiving traffic")

	flag.Parse()

	routes, body, code, err := ParseRoutes(*cfgFile)
	if err != nil {
		log.Fatalf("config loading failed: %s", err)
	}

	cfg = config{
		port:       *port,
		routes:     routes,
		body:       body,
		statusCode: code,
	}
}
