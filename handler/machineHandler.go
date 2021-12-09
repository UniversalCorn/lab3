package hendler

import (
	"fmt"
	"net/http"
)

func getMachines(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf(fmt.Sprintf("Here will get machines\n"))
}

func addMachine(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Here will add machine\n")
}
