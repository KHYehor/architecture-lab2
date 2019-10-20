//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/KHYehor/architecture-lab2/server/tablets"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*ChatApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from channels package.
		tablets.Providers,
		// Provide ChatApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(ChatApiServer), "Port", "TabletsHandler"),
	)
	return nil, nil
}
