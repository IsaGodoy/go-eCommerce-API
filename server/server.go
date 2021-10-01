package server

import (
	"net/http"

	"github.com/IsaGodoy/go-eCommerce-API/routes"
)

type Server struct {
	port   string
	router *routes.Router
}

func NewServer(_port string) *Server {
	return &Server{
		port:   _port,
		router: routes.NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	return err
}

func (s *Server) AddRoute(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.Rules[path]

	if !exist {
		s.router.Rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.Rules[path][method] = handler
}
