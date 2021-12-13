package uniqueStore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UniqueStore struct {
	MStore *MachineStore
	DStore *DisckStore
}

func NewStore(db *sql.DB) *UniqueStore {
	dstore := NewDickStore(db)
	mstore := NewMachineStore(db)
	return &UniqueStore{MStore: mstore, DStore: dstore}
}

func (gs *UniqueStore) GetMachine() ([]*Machine, error) {
	return gs.MStore.GetMachine()
}
