package corona

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
)

//
type notification struct {
}

//
func HandlerNotifications(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("")
	if err != nil {
		fmt.Print(err.Error())
	}
}
