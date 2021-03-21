package corona

import (
	"encoding/json"
	"fmt"
	"main/functions"
	"net/http"
	"strings"
)

//
type covidCases struct {
	Country               string `json:"country"`
	Continent             string `json:"continent"`
	Scope                 string `json:"scope"`
	Confirmed             string `json:"confirmed"`
	Recovered             string `json:"recovered"`
	Population_percentage string `json:"population_percentage"`
}

//
func HandlerCoronaCase(w http.ResponseWriter, r *http.Request) {
	var err error
	// 0local , 1corona , 2v1 , 3country , 4countryNameandSCOPE
	urlArray := strings.Split(functions.GetURL(r), "/")

	if len(urlArray) != 5 {
		fmt.Print(err.Error())
		return
	}

	// "Norway?scope=2020-12-01-2021-01-31"
	// Norway, scope=2020-12-01-2021-01-31
	urlParameters := strings.Split(urlArray[4], "?")
	fmt.Print(urlParameters)

	if len(urlParameters) == 2 {
		fmt.Print(len(urlParameters))
		//fmt.Print(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode("")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

//
/* func getCoronaCases(e *covidCases, startDate string, endDate string, scope string) error {

} */
