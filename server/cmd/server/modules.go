//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/pavlovskyive/architecture-lab-3/server/menu"
	"github.com/pavlovskyive/architecture-lab-3/server/orders"
)

// ComposeAPIServer will create an instance of RestaurantAPIServer according to providers defined in this file
func ComposeAPIServer(port HTTPPortNumber) (*RestaurantAPIServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from menu package.
		menu.Providers,
		// Add providers from orders package
		orders.Providers,
		// Provide RestaurantAPIServer instantiating the structure and injecting menu and orders handler and port number
		wire.Struct(new(RestaurantAPIServer), "Port", "MenuHandler", "OrdersHandler"),
	)
	return nil, nil
}
