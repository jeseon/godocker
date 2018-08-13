package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{}
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	s.Router.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nDisallow: /"))
	})
	return s
}

func (s *Server) Run() {
	http.ListenAndServe(":5000", s)
}