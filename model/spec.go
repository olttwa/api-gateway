package model

type RouteSpec struct {
	PathPrefix string `yaml:"path_prefix"`
	Backend    string `yaml:"backend"`
}

type DefaultResponse struct {
	Body       string `yaml:"body"`
	StatusCode int    `yaml:"status_code"`
}

type Backend struct {
	Name        string   `yaml:"name"`
	MatchLabels []string `yaml:"match_labels"`
}

type Spec struct {
	Routes          []RouteSpec     `yaml:"routes"`
	DefaultResponse DefaultResponse `yaml:"default_response"`
	Backends        []Backend       `yaml:"backends"`
}
