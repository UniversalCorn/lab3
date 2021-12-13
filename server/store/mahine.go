package uniqueStore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type MachineStore struct {
	Db *sql.DB
}

type Machine struct {
	Id             int    `json:"Id"`
	MachineName    string `json:"MachineName"`
	CpuCount       int    `json:"CpuCount"`
	TotalDiskSpace int    `json:"TotalDiskSpace"`
}

func NewMachineStore(db *sql.DB) *MachineStore {
	return &MachineStore{Db: db}
}

func (m *MachineStore) GetMachine() ([]*Machine, error) {
	rows, err := m.Db.Query("SELECT * FROM machines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Machine
	for rows.Next() {
		var c Machine
		if err := rows.Scan(&c.Id, &c.MachineName, &c.CpuCount, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		fmt.Print(c)
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Machine, 0)
	}
	return res, nil
}
