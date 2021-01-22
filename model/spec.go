package model

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

type Spec struct {
	Routes          []route         `yaml:"routes"`
	DefaultResponse defaultResponse `yaml:"default_response"`
	Backends        []backend       `yaml:"backends"`
}
