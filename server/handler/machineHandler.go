package hendler

import (
	"net/http"

	gs "github.com/IP94-rocketBunny-architecture/lab3/server/store"
	"github.com/IP94-rocketBunny-architecture/lab3/tools"
	_ "github.com/go-sql-driver/mysql"
)

func getMachines(db *gs.UniqueStore, rw http.ResponseWriter, req *http.Request) {
	if res, err := db.GetMachine(); err != nil {
		tools.WriteJsonBadRequest(rw, err.Error())
	} else {
		tools.WriteJsonOk(rw, res)
	}
}
