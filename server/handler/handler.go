package hendler

import (
	"net/http"

	gs "github.com/UniversalCorn/lab3/server/store"
)

type Handlers struct {
	db *gs.UniqueStore
}

func (h *Handlers) HandleMachines(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getMachines(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) HandleDisks(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getDisks(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) HandleVolume(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		increaseMachineSpace(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) HandleFindByIdDisk(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		findByIdDisk(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) HandleFindByIdMachine(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		findByIdMachine(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}
