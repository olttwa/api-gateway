package handler_test

import (
	"net/http"
	"net/http/httptest"
	"rgate/handler"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDefault(t *testing.T) {
	body := "default response"
	code := http.StatusForbidden

	d := handler.Default([]byte(body), code)

	recorder := httptest.NewRecorder()
	d.ServeHTTP(recorder, nil)

	assert.Equal(t, body, recorder.Body.String())
	assert.Equal(t, code, recorder.Code)
}
