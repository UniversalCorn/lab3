package uniqueStore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UniqueStore struct {
	MStore *MachineStore
	DStore *DiskStore
}

func NewStore(db *sql.DB) *UniqueStore {
	dstore := NewDiskStore(db)
	mstore := NewMachineStore(db)
	return &UniqueStore{MStore: mstore, DStore: dstore}
}

func (gs *UniqueStore) GetMachines() ([]*Machine, error) {
	return gs.MStore.GetMachines()
}

func (gs *UniqueStore) GetDisks() ([]*Disk, error) {
	return gs.DStore.GetDisks()
}

func (gs *UniqueStore) FindByIdDisk(Id int) ([]*Disk, error) {
	return gs.DStore.FindByIdDisk(Id)
}

func (gs *UniqueStore) IncreaseMachineSpace(idArr []int) ([]*Machine, error) {
	res, err := gs.DStore.FindByIdDisk(idArr[1])
	if err != nil {
		return nil, err
	}
	result, err := gs.MStore.IncreaseMachineSpace(idArr[0], res)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (gs *UniqueStore) FindByIdMachine(Id int) ([]*Machine, error) {
	return gs.MStore.FindByIdMachine(Id)
}
