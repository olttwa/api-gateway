package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type route struct {
	PathPrefix string `yaml:"path_prefix"`
	Backend    string `yaml:"backend"`
}

type defaultResponse struct {
	Body       string `yaml:"body"`
	StatusCode int    `yaml:"status_code"`
}

type backend struct {
	Name        string   `yaml:"name"`
	MatchLabels []string `yaml:"match_labels"`
}

type spec struct {
	Routes          []route         `yaml:"routes"`
	DefaultResponse defaultResponse `yaml:"default_response"`
	Backends        []backend       `yaml:"backends"`
}

var ReadFile = ioutil.ReadFile

func ParseRoutes(file string) ([]Route, []byte, int, error) {
	y, err := ReadFile(file)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("error reading file: %w", err)
	}

	var s spec
	err = yaml.Unmarshal(y, &s)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("error unmarshaling yaml: %w", err)
	}

	backends := make(map[string][]string, len(s.Backends))
	for _, b := range s.Backends {
		backends[b.Name] = b.MatchLabels
	}

	routes := make([]Route, 0, len(s.Routes))
	for _, r := range s.Routes {
		ml, ok := backends[r.Backend]
		if !ok {
			return nil, nil, 0, fmt.Errorf("backend not defined: %s", r.Backend)
		}

		route := Route{
			PathPrefix:  r.PathPrefix,
			MatchLabels: ml,
		}
		routes = append(routes, route)
	}

	return routes, []byte(s.DefaultResponse.Body), s.DefaultResponse.StatusCode, nil
}
