package main

import (
	"database/sql"
	"log"

	"github.com/IP94-rocketBunny-architecture/lab3/db"
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
	// db, err := NewDbConnection()
	// if err != nil {
	// 	log.Fatalf("failed to initialize db: %s", err.Error())
	// }
	server := new(Server)
	if err := server.Run("8000"); err != nil {
		log.Fatal("Cannot initialize forums server: %s", err.Error())
	}
}
