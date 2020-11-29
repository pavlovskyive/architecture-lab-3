package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pavlovskyive/architecture-lab-3/server/menu"
	"github.com/pavlovskyive/architecture-lab-3/server/orders"
)

// HTTPPortNumber stores the number of port
type HTTPPortNumber int

// RestaurantAPIServer configures necessary handlers and starts listening on a configured port
type RestaurantAPIServer struct {
	Port          HTTPPortNumber
	MenuHandler   menu.HTTPHandlerFunc
	OrdersHandler orders.HTTPHandlerFunc
	server        *http.Server
}

// Start will set all handlers and start listening
// If this methods succeeds, it does not return until server is shut down
// Returned error will never be nil
func (s *RestaurantAPIServer) Start() error {
	if s.MenuHandler == nil {
		return fmt.Errorf("menu HTTP handler is not defined - cannot start")
	}
	if s.OrdersHandler == nil {
		return fmt.Errorf("orders HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/menu", s.MenuHandler)
	handler.HandleFunc("/orders", s.OrdersHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stop will shut down previously started HTTP server
func (s *RestaurantAPIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
