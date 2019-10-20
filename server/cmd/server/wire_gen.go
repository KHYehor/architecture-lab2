// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/KHYehor/architecture-lab2/server/tablets"
)

// Injectors from modules.go:

func ComposeApiServer(port HttpPortNumber) (*TabletApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := tablets.NewStore(db)
	httpHandlerFunc := tablets.HttpHandler(store)
	tabletApiServer := &TabletApiServer{
		Port:           port,
		TabletsHandler: httpHandlerFunc,
	}
	return tabletApiServer, nil
}
