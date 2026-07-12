package server

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {

	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Mux() *http.ServeMux {
	return s.mux
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.mux)
}
