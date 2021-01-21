package utils

import "net/http"

type ResponseRecorder struct {
	http.ResponseWriter
	status int
}

// Overwrite WriteHeader method to capture status.
func (r *ResponseRecorder) WriteHeader(s int) {
	r.ResponseWriter.WriteHeader(s)
	r.status = s
}

func (r *ResponseRecorder) Status() int {
	return r.status
}
