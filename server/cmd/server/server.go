package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jn-lp/se-lab3/server/plants"
)

type HttpPortNumber int

// PlantsApiServer configures necessary handlers and starts listening on a configured port.
type PlantsApiServer struct {
	Port HttpPortNumber

	PlantsHandler plants.HTTPHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *PlantsApiServer) Start() error {
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
func (s *PlantsApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
