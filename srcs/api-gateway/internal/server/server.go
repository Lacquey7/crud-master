package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Addr    string
	Port    string
	Handler http.Handler
}

func NewServer(addr, port string, handler http.Handler) *Server {
	return &Server{
		Addr:    addr,
		Port:    port,
		Handler: handler,
	}
}

func (s *Server) Start() {
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", s.Addr, s.Port),
		Handler:        s.Handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
