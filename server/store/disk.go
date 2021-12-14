package uniqueStore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DiskStore struct {
	Db *sql.DB
}

type Disk struct {
	Id             int `json:"Id"`
	TotalDiskSpace int `json:"TotalDiskSpace"`
}

func NewDiskStore(db *sql.DB) *DiskStore {
	return &DiskStore{Db: db}
}

func (m *DiskStore) GetDisks() ([]*Disk, error) {
	rows, err := m.Db.Query("SELECT * FROM discks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Disk
	for rows.Next() {
		var c Disk
		if err := rows.Scan(&c.Id, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		fmt.Print(c)
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Disk, 0)
	}
	return res, nil
}

func (m *DiskStore) FindByIdDisk(Id int) ([]*Disk, error) {
	rowsM, err := m.Db.Query(`SELECT * FROM discks WHERE discks.Id = $1`, Id)
	if err != nil {
		return nil, err
	}
	defer rowsM.Close()

	var res []*Disk
	for rowsM.Next() {
		var c Disk
		if err := rowsM.Scan(&c.Id, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}

	if res == nil {
		res = make([]*Disk, 0)
	}

	return res, nil
}
