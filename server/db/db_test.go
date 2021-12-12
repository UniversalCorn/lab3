package db

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "lab3",
		User:       "mysql",
		Password:   "mysql",
		Host:       "localhost",
		DisableSSL: true,
	}
	fmt.Println(conn.ConnectionURL())
	sql.Open("mysql", conn.ConnectionURL())
	if conn.ConnectionURL() != "mysql://mysql:mysql@localhost/tables?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
