package corona

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
)

//
type covidCases struct {
}

//
func HandlerCoronaCases(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("")
	if err != nil {
		fmt.Print(err.Error())
	}
}

//
func getCoronaCases(e *covidCases, startDate string, endDate string, scope string) error {
	var err error
	return err
}
