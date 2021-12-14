package main

import (
	"context"
	"fmt"
	"net/http"

	h "github.com/UniversalCorn/lab3/server/handler"
)

type Server struct {
	httpServer *http.Server
	Handlers   *h.Handlers
	Senv       *ServerEnv
}

func (ser *Server) Run() error {
	handlersCollection := map[string]http.HandlerFunc{
		"/machines": ser.Handlers.HandleMachines,
		"/discks":   ser.Handlers.HandleDiscks,
		// "/update":   ser.Handlers.HandleUpdate,
	}
	for route, handler := range handlersCollection {
		http.Handle(route, handler)
	}
	runnable := ser.Senv.Host + ":" + fmt.Sprint(ser.Senv.Port)
	ser.httpServer = &http.Server{Addr: runnable}
	fmt.Printf("Server is running on port: %d, host: %s\n", ser.Senv.Port, ser.Senv.Host)
	return ser.httpServer.ListenAndServe()
}

func (ser *Server) Close() error {
	if ser.httpServer == nil {
		return fmt.Errorf("Server was not started")
	}
	return ser.httpServer.Shutdown(context.Background())
}
