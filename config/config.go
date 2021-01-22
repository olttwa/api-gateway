package config

import (
	"flag"
	"io/ioutil"
	"log"
	"rgate/model"
	"rgate/utils"

	"gopkg.in/yaml.v2"
)

type Config struct {
	port        string
	defaultBody []byte
	defaultCode int
	routes      []model.Route
}

func (cfg *Config) parseRoutes(file string) {
	y, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	var s model.Spec
	err = yaml.Unmarshal(y, &s)
	if err != nil {
		log.Fatalf("error unmarshaling yaml: %s", err)
	}

	cfg.defaultBody = []byte(s.DefaultResponse.Body)
	cfg.defaultCode = s.DefaultResponse.StatusCode
	cfg.routes, err = utils.MatchRoutesToBackend(s)
	if err != nil {
		log.Fatalf("error matching routes to backends: %s", err)
	}
}

func (cfg *Config) Port() string {
	return cfg.port
}

func (cfg *Config) Routes() []model.Route {
	return cfg.routes
}

func (cfg *Config) DefaultBody() []byte {
	return cfg.defaultBody
}

func (cfg *Config) DefaultCode() int {
	return cfg.defaultCode
}

func Load() *Config {
	cfgFile := flag.String("config", "config.yml", "config filename")
	port := flag.String("port", "8080", "port for receiving traffic")
	flag.Parse()

	c := Config{port: *port}
	c.parseRoutes(*cfgFile)
	return &c
}
