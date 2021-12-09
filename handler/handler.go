package hendler

import (
	"net/http"

	gs "github.com/IP94-rocketBunny-architecture/lab3/genStore"
)

type Handlers struct {
	db *gs.GenSore
}

func (h *Handlers) HandleMachines(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getMachines(rw, req)
	} else if req.Method == "POST" {
		addMachine(rw, req)
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
