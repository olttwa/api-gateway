package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"rgate/middleware"
	"testing"

	"gotest.tools/v3/assert"
)

func TestStatsMW(t *testing.T) {
	h := func(code int) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		})
	}

	recorder := httptest.NewRecorder()
	s := middleware.StatsMW()
	s.Middleware(h(100)).ServeHTTP(recorder, nil)
	s.Middleware(h(200)).ServeHTTP(recorder, nil)
	s.Middleware(h(399)).ServeHTTP(recorder, nil)
	s.Middleware(h(400)).ServeHTTP(recorder, nil)
	s.Middleware(h(503)).ServeHTTP(recorder, nil)

	assert.Equal(t, 2, s.Success())
	assert.Equal(t, 2, s.Errors())
}
