package plants

import "github.com/google/wire"

// Providers ist of providers for plants components.
var Providers = wire.NewSet(NewStore, HTTPHandler)
