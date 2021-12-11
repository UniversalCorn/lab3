package uniqueStore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DisckStore struct {
	Db *sql.DB
}

func NewDickStore(db *sql.DB) *DisckStore {
	return &DisckStore{Db: db}
}
