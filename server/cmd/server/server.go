package main

import (
	"context"
	"net/http"

	h "github.com/IP94-rocketBunny-architecture/lab3/handler"
)

type Server struct {
	httpServer *http.Server
	Handlers   *h.Handlers
}

func (ser *Server) Run(port string) error {
	handlersCollection := map[string]http.HandlerFunc{
		"/machines": ser.Handlers.HandleMachines,
		"/discks":   ser.Handlers.HandleDiscks,
		"/machine":  ser.Handlers.GetMachine,
		"/disck":    ser.Handlers.GetDisck,
	}
	for route, handler := range handlersCollection {
		http.Handle(route, handler)
	}
	ser.httpServer = &http.Server{
		Addr: ":" + port,
	}
	return ser.httpServer.ListenAndServe()
}

func (ser *Server) Shutdown(ctx context.Context) error {
	return ser.httpServer.Shutdown(ctx)
}
