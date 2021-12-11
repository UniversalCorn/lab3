// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/IP94-rocketBunny-architecture/lab3/server/handler"
	"github.com/IP94-rocketBunny-architecture/lab3/server/store"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewServer(senv *ServerEnv) (*Server, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	uniqueStoreUniqueStore := uniqueStore.NewStore(db)
	handlers := hendler.NewHandler(uniqueStoreUniqueStore)
	server := &Server{
		Handlers: handlers,
		Senv:     senv,
	}
	return server, nil
}

// wire.go:

var providers = wire.NewSet(uniqueStore.NewStore, hendler.NewHandler)
