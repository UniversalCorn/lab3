package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/UniversalCorn/lab3/server/db"
)

type ServerEnv struct {
	Port int
	Host string
}

var (
	servPORT = flag.Int("p", 8000, "HTTP port number")
	servHOST = flag.String("h", "localhost", "HTTP host name")
)

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "lab3",
		User:       "mysql",
		Password:   "mysql",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	flag.Parse()
	senv := &ServerEnv{Port: *servPORT, Host: *servHOST}
	server, err := NewServer(senv)
	if err != nil {
		log.Fatalf("Cannot initialize forums server: %s", err)
	}

	go func() {
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel
		if err := server.Close(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	}()

	if server.Run() == http.ErrServerClosed {
		log.Printf("HTTP server stopped")
	} else {
		log.Fatalf("Cannot start HTTP server: %s", err)
	}
}
