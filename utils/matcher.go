package utils

import (
	"fmt"
	"rgate/model"
)

func MatchRoutesToBackend(s model.Spec) ([]model.Route, error) {
	backends := make(map[string][]string, len(s.Backends))
	for _, b := range s.Backends {
		backends[b.Name] = b.MatchLabels
	}

	routes := make([]model.Route, 0, len(s.Routes))
	for _, r := range s.Routes {
		ml, ok := backends[r.Backend]
		if !ok {
			return nil, fmt.Errorf("backend not defined for route: %s", r.PathPrefix)
		}

		route := model.Route{
			PathPrefix:  r.PathPrefix,
			MatchLabels: ml,
		}
		routes = append(routes, route)
	}
	return routes, nil
}
