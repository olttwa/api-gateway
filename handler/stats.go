package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rgate/model"
	"rgate/utils"
)

type Stats interface {
	Success() int
	Errors() int
	Mean() int
	Percentile(float64) int
}

type stats struct {
	s Stats
}

func (s stats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := model.Stats{
		Requests: model.Requests{
			Success: s.s.Success(),
			Error:   s.s.Errors(),
		},
		Latency: model.Latency{
			Average: s.s.Mean(),
			P95:     s.s.Percentile(95),
			P99:     s.s.Percentile(99),
		},
	}

	data, err := json.Marshal(d)
	if err != nil {
		utils.ErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(data); err != nil {
		log.Printf("error writing stats response: %s\n", err)
	}
}

func StatsHandler(s Stats) http.Handler {
	return stats{s}
}
