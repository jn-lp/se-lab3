package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jn-lp/se-lab3/server/plants"
)

// HTTPPortNumber is a type for server's HTTP port
type HTTPPortNumber int

// PlantsAPIServer configures necessary handlers and starts listening on a configured port.
type PlantsAPIServer struct {
	Port HTTPPortNumber

	PlantsHandler plants.HTTPHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *PlantsAPIServer) Start() error {
	if s.PlantsHandler == nil {
		return fmt.Errorf("plants HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/plants", s.PlantsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stop will shut down previously started HTTP server.
func (s *PlantsAPIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
