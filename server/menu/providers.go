package menu

import "github.com/google/wire"

// Providers is a set of providers for menu components.
var Providers = wire.NewSet(NewStore, HTTPHandler)
