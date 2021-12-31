package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

// New initializes and returns a new server
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}
	return &server, nil
}

// Close server resources
func (s *Server) Close() error {
	// TODO: close server resources
	return nil
}

// Start server
func (s *Server) Start() {
	log.Printf("Server running on http://localhost%s", s.server.Addr)
	log.Fatal(s.server.ListenAndServe())
}
