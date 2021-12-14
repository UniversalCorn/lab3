package uniqueStore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DisckStore struct {
	Db *sql.DB
}

type Disck struct {
	Id             int `json:"Id"`
	TotalDiskSpace int `json:"TotalDiskSpace"`
}

func NewDickStore(db *sql.DB) *DisckStore {
	return &DisckStore{Db: db}
}

func (m *DisckStore) GetDicks() ([]*Disck, error) {
	rows, err := m.Db.Query("SELECT * FROM discks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Disck
	for rows.Next() {
		var c Disck
		if err := rows.Scan(&c.Id, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		fmt.Print(c)
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Disck, 0)
	}
	return res, nil
}
