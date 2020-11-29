package orders

import "github.com/google/wire"

// Providers is a set of providers for orders components
var Providers = wire.NewSet(NewStore, HTTPHandler)
