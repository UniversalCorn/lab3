//go:build wireinject
// +build wireinject

package main

import (
	h "github.com/IP94-rocketBunny-architecture/lab3/server/handler"
	gs "github.com/IP94-rocketBunny-architecture/lab3/server/store"
	"github.com/google/wire"
)

var providers = wire.NewSet(gs.NewStore, h.NewHandler)

func NewServer(senv *ServerEnv) (*Server, error) {
	wire.Build(
		NewDbConnection,
		providers,
		wire.Struct(new(Server), "Handlers", "Senv"),
	)

	return nil, nil
}
