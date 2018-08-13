package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

var reqs = []struct {
	path string
	body string
	code int
}{
	{"/favicon.ico", "", http.StatusOK},
	{"/robots.txt", "User-agent: *\nDisallow: /", http.StatusOK},
}

func TestServer(t *testing.T) {
	s := NewServer()

	for _, req := range reqs {
		r, _ := http.NewRequest("GET", req.path, nil)
		w := httptest.NewRecorder()
		s.Router.ServeHTTP(w, r)

		assert.Equal(t, req.body, w.Body.String(), "Should be the same.")
		assert.Equal(t, req.code, w.Code, "Should be the same.")
	}
}
