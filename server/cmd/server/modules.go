//+build wireinject

package main

import (
	"github.com/KHYehor/architecture-lab2/server/tablets"
	"github.com/google/wire"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*TabletApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from channels package.
		tablets.Providers,
		// Provide ChatApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(TabletApiServer), "Port", "TabletsHandler"),
	)
	return nil, nil
}
