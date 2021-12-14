package hendler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	gs "github.com/UniversalCorn/lab3/server/store"
	"github.com/UniversalCorn/lab3/tools"
	_ "github.com/lib/pq"
)

type MachineID struct {
	MiD int `json:"Id"`
}

func getMachines(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.GetMachines(); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}

func increaseMachineSpace(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resId := []MachineID{}

	jsonString, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	if err := json.Unmarshal([]byte(jsonString), &resId); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}

	if len(resId) == 0 {
		tools.WriteJsonBadRequest(rw, "machine ID and disck ID is not provided")
		return
	}

	if resId[0].MiD < 1 && resId[1].MiD < 1 {
		tools.WriteJsonBadRequest(rw, "machine ID or disck ID is not provided")
		return
	}

	if res, err := db.IncreaseMachineSpace([]int{resId[0].MiD, resId[1].MiD}); err != nil {
		tools.WriteJsonInternalError(rw, "machine with this ID is not exists")
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
