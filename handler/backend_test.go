package handler_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"rgate/handler"
	"rgate/mocks"
	"rgate/model"
	"rgate/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBackendReverseProxy(t *testing.T) {
	method := "POST"
	path := "/test/reverse/proxy"
	authHeader := "Bearer: 12345"
	body := []byte("request body")

	mockHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		assert.Equal(t, method, r.Method)
		assert.Equal(t, path, r.URL.Path)
		assert.Equal(t, authHeader, r.Header.Get("Authorization"))
		assert.NotNil(t, r.Header.Get("X-Forwarded-For"))

		rBody, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, body, rBody)
	})

	reverseProxy := httptest.NewServer(mockHandler)
	defer reverseProxy.Close()
	reverseProxyURL, err := url.Parse(reverseProxy.URL)
	assert.NoError(t, err)

	d := &mocks.Docker{}
	c := []model.Container{{IP: reverseProxyURL.Hostname(), Port: reverseProxyURL.Port()}}
	d.On("ListContainers", mock.Anything, []string{"foo=bar"}).Return(c, nil)
	b := handler.Backend(d, []string{"foo=bar"})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set("Authorization", authHeader)
	b.ServeHTTP(recorder, req)
}

func TestBackendError(t *testing.T) {
	d := &mocks.Docker{}
	d.On("ListContainers", mock.Anything, mock.Anything).Return([]model.Container{}, nil)
	b := handler.Backend(d, []string{})

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	b.ServeHTTP(recorder, req)

	assert.Equal(t, utils.BackendErrorCode, recorder.Code)
}
