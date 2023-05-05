package http

import (
	"ServerProtocols/net"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	net.ServerConfig
}

func New(configs ...net.ServerConfigFunc) *Server {
	config := net.DefaultConfig()
	for _, fn := range configs {
		fn(&config)
	}
	return &Server{
		ServerConfig: config,
	}
}

func (s *Server) Handle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

func (s *Server) Listen() error {
	addr := fmt.Sprintf("%s:%s", s.ServerConfig.Host, s.ServerConfig.Port)
	log.Printf("Server serve at %s", addr)
	return http.ListenAndServe(addr, nil)
}
