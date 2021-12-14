package hendler

import (
	"net/http"

	gs "github.com/UniversalCorn/lab3/server/store"
	"github.com/UniversalCorn/lab3/tools"
)

func getDiscks(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.GetDicks(); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
