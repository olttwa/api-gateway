package model

type Requests struct {
	Success int `json:"success"`
	Error   int `json:"error"`
}

type Latency struct {
	Average int `json:"average"`
	P95     int `json:"p95"`
	P99     int `json:"p99"`
}

type Stats struct {
	Requests Requests `json:"requests_count"`
	Latency  Latency  `json:"latency_ms"`
}
