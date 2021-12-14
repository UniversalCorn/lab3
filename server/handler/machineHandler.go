package hendler

import (
	"encoding/json"
	"log"
	"net/http"

	gs "github.com/UniversalCorn/lab3/server/store"
	"github.com/UniversalCorn/lab3/tools"
	_ "github.com/lib/pq"
)

type ResponseName struct {
	MachineID int `json:"Id"`
	DisckId   int `json:"Id"`
}

func getMachines(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.GetMachine(); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}

func updateMachineById(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var resName ResponseName

	if err := json.NewDecoder(req.Body).Decode(&resName); err != nil {
		log.Printf("Error decoding interest input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	if resName.MachineID < 1 || resName.DisckId < 1 {
		tools.WriteJsonBadRequest(rw, "machine ID or disck ID is not provided")
		return
	}

	if res, err := db.FindById(resName.MachineID, resName.DisckId); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
