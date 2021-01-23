package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rgate/model"
	"rgate/utils"
)

type stats interface {
	Success() int
	Errors() int
	Mean() int
	Percentile(float64) int
}

type statsHandler struct {
	stats
}

func (h statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := model.Stats{
		Requests: model.Requests{
			Success: h.Success(),
			Error:   h.Errors(),
		},
		Latency: model.Latency{
			Average: h.Mean(),
			P95:     h.Percentile(95),
			P99:     h.Percentile(99),
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

func Stats(s stats) http.Handler {
	return statsHandler{s}
}
