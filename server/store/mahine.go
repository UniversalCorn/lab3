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

func (m *MachineStore) GetMachines() ([]*Machine, error) {
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

func (m *MachineStore) IncreaseMachineSpace(MiD int, disck []*Disk) ([]*Machine, error) {
	rowsM, err := m.Db.Query(`SELECT * FROM machines WHERE machines.Id = $1`, MiD)
	if err != nil {
		return nil, err
	}
	defer rowsM.Close()

	var res []*Machine
	for rowsM.Next() {
		var c Machine
		if err := rowsM.Scan(&c.Id, &c.MachineName, &c.CpuCount, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		dT := disck[len(disck)-1].TotalDiskSpace
		c.TotalDiskSpace += dT
		res = append(res, &c)
	}

	if res == nil {
		res = make([]*Machine, 0)
	}

	return res, nil
}

func (m *MachineStore) FindByIdMachine(MiD int) ([]*Machine, error) {
	rowsM, err := m.Db.Query(`SELECT * FROM machines WHERE machines.Id = $1`, MiD)
	if err != nil {
		return nil, err
	}
	defer rowsM.Close()

	var res []*Machine
	for rowsM.Next() {
		var c Machine
		if err := rowsM.Scan(&c.Id, &c.MachineName, &c.CpuCount, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}

	if res == nil {
		res = make([]*Machine, 0)
	}

	return res, nil
}