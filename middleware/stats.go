package middleware

import (
	"log"
	"math"
	"net/http"
	"rgate/utils"
	"sync"
	"time"

	"github.com/montanaflynn/stats"
)

const (
	successThreshold = 200
	errorThreshold   = 399
)

type Stats struct {
	m             *sync.Mutex
	success       int
	errors        int
	responseTimes []float64
}

func (s *Stats) Success() int {
	return s.success
}

func (s *Stats) Errors() int {
	return s.errors
}

func (s *Stats) Mean() int {
	m, err := stats.Mean(s.responseTimes)
	if err != nil {
		log.Printf("error finding mean response time: %s", err)
		return 0
	}
	return int(math.Round(m))
}

func (s *Stats) Percentile(p float64) int {
	m, err := stats.Percentile(s.responseTimes, p)
	if err != nil {
		log.Printf("error finding percentile response time: %s", err)
		return 0
	}
	return int(math.Round(m))
}

func (s *Stats) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &utils.ResponseRecorder{ResponseWriter: w}
		start := time.Now()

		next.ServeHTTP(recorder, r)

		responseTime := time.Since(start)

		// Handle race conditions.
		s.m.Lock()
		defer s.m.Unlock()

		if recorder.Status() > errorThreshold {
			s.errors++
		} else if recorder.Status() >= successThreshold {
			s.success++
		}

		s.responseTimes = append(s.responseTimes, float64(responseTime.Milliseconds()))
	})
}

func StatsMw() *Stats {
	return &Stats{
		m: &sync.Mutex{},
	}
}
