package corona

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//diagnosis struct
type diagnosis struct {
	Mmediagroupapi  int    `json:"mmediagroupapi"`
	Covidtrackerapi int    `json:"covidtrackerapi"`
	Registered      int    `json:"registered"`
	Version         string `json:"version"`
	Uptime          string `json:"uptime"`
}

//variable to keep track of uptime
var StartTime time.Time

//calculates uptime
func getUptime() float64 {
	return time.Since(StartTime).Seconds()
}

//handles diagnosis
func HandlerDiagnosis(w http.ResponseWriter, r *http.Request) {
	var diag diagnosis
	var err error

	//get mmediagroup API status code
	resp, err := http.Get("https://covid-api.mmediagroup.fr/v1/history?country=Norway&status=Confirmed")
	if err != nil {
		fmt.Print(err.Error())
	}
	//insert mmediagroup status code to diag struct
	diag.Mmediagroupapi = resp.StatusCode

	//get covidtracker API status code
	resp, err = http.Get("https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/date-range/2021-03-21/2021-03-21")
	if err != nil {
		fmt.Print(err.Error())
	}
	//insert covidtracker status code to diag struct
	diag.Covidtrackerapi = resp.StatusCode
	//Registered webhooks
	diag.Registered = 0
	//Version
	diag.Version = "v1"
	//Uptime
	diag.Uptime = fmt.Sprintf("%f", getUptime())

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(diag)
	if err != nil {
		fmt.Print(err.Error())
	}
}
