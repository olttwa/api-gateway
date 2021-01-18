package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type route struct {
	PathPrefix string `yaml:"path_prefix"`
	Backend    string `yaml:"backend"`
}

type defaultResponse struct {
	Body       string `yaml:"body"`
	StatusCode string `yaml:"status_code"`
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

func ParseRoutes(file string) ([]Route, error) {
	y, err := ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var s spec
	err = yaml.Unmarshal(y, &s)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling yaml: %w", err)
	}

	backends := make(map[string][]string, len(s.Backends))
	for _, b := range s.Backends {
		backends[b.Name] = b.MatchLabels
	}

	routes := make([]Route, 0, len(s.Routes))
	for _, r := range s.Routes {
		labels, ok := backends[r.Backend]
		if !ok {
			return nil, fmt.Errorf("backend not defined: %s", r.Backend)
		}

		ml := make(map[string]string, len(labels))
		for _, l := range labels {
			split := strings.Split(l, "=")
			ml[split[0]] = split[1]
		}

		route := Route{
			PathPrefix:  r.PathPrefix,
			MatchLabels: ml,
		}
		routes = append(routes, route)
	}

	return routes, nil
}
