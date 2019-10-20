package main

import (
	"context"
	"fmt"
	"github.com/KHYehor/architecture-lab2/server/tablets"
	"net/http"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type TabletApiServer struct {
	Port HttpPortNumber

	TabletsHandler tablets.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *TabletApiServer) Start() error {
	if s.TabletsHandler == nil {
		return fmt.Errorf("Tablets HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/sendData", s.TabletsHandler)
	handler.HandleFunc("/getData", s.TabletsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *TabletApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
