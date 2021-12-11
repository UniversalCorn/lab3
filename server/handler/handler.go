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
		getDiscks(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handlers) GetMachine(rw http.ResponseWriter, req *http.Request) {
	//отадть юзеру все машины
}

func (h *Handlers) GetDisck(rw http.ResponseWriter, req *http.Request) {
	//отадть юзеру все диски
}
