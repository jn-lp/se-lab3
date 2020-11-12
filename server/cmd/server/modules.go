//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jn-lp/se-lab3/server/plants"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*PlantsApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from plants package.
		plants.Providers,
		// Provide ChatApiServer instantiating the structure and injecting plants handler and port number.
		wire.Struct(new(PlantsApiServer), "Port", "PlantsHandler"),
	)
	return nil, nil
}
