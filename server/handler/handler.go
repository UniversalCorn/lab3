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

func (h *Handlers) HandleDiscks(rw http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		getDiscks(h.db, rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) HandleUpdate(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		updateMachineById(h.db, rw, req)
	}
}
