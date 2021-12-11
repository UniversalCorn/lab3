package hendler

import (
	gs "github.com/UniversalCorn/lab3/server/store"
)

func NewHandler(gs *gs.UniqueStore) *Handlers {
	return &Handlers{gs}
}
