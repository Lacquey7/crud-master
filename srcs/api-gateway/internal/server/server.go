package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Address string
	Port    string
	Router  http.Handler
}

// New crée une nouvelle instance de Server avec les paramètres fournis
func New(address, port string, router http.Handler) *Server {
	return &Server{
		Address: address,
		Port:    port,
		Router:  router,
	}
}

// Run lance le serveur avec des paramètres avancés
func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%s", s.Address, s.Port)

	// Configuration avancée du serveur HTTP
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           s.Router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 Mo
	}

	log.Printf("Démarrage du serveur sur http://%s", addr)
	return httpServer.ListenAndServe()
}
