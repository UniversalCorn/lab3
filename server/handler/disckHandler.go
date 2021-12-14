package hendler

import (
	"encoding/json"
	"net/http"

	gs "github.com/UniversalCorn/lab3/server/store"
	"github.com/UniversalCorn/lab3/tools"
)

type DisckId struct {
	ID int `json:"Id"`
}

func getDisks(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.GetDisks(); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
	} else {
		tools.WriteJsonOk(rw, res)
	}
}

func findByIdDisk(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var resId DisckId

	if err := json.NewDecoder(req.Body).Decode(&resId); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}

	if resId.ID < 1 {
		tools.WriteJsonBadRequest(rw, "disck ID is not provided")
		return
	}

	if res, err := db.FindByIdDisk(resId.ID); err != nil {
		tools.WriteJsonInternalError(rw, "disk with this ID is not exists")
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
