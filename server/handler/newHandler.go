package hendler

import (
	gs "github.com/IP94-rocketBunny-architecture/lab3/server/store"
)

func NewHandler(gs *gs.UniqueStore) *Handlers {
	return &Handlers{gs}
}
