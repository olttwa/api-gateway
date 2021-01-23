package handler_test

import (
	"net/http/httptest"
	"rgate/handler"
	"rgate/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsHandler(t *testing.T) {
	s := &mocks.Stats{}
	h := handler.Stats(s)

	s.On("Success").Return(7)
	s.On("Errors").Return(4)
	s.On("Mean").Return(3)
	s.On("Percentile", float64(95)).Return(9)
	s.On("Percentile", float64(99)).Return(13)

	recorder := httptest.NewRecorder()
	h.ServeHTTP(recorder, nil)

	expectedResponse := `
	{
		"requests_count": {
			"success": 7,
			"error": 4
		},
		"latency_ms": {
			"average": 3,
			"p95": 9,
			"p99": 13
		}
	}`
	assert.JSONEq(t, expectedResponse, recorder.Body.String())
}
