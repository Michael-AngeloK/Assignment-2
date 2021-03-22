package corona

import (
	"encoding/json"
	"fmt"
	"main/functions"
	"math"
	"net/http"
	"net/url"
	"strings"
)

// NB!: NEEDS DATE VERIFICATION, IF SCOPE EMPTY AND CAN REDO UPDATEOUTPUT FUNCTION, ERRORS
//
type coronaCases struct {
	All struct {
		Country    string         `json:"country"`
		Population int            `json:"population"`
		Continent  string         `json:"continent"`
		Dates      map[string]int `json:"dates"`
	} `json:"All"`
}

//
type outputCoronaCases struct {
	Country               string  `json:"country"`
	Continent             string  `json:"continent"`
	Scope                 string  `json:"scope"`
	Confirmed             int     `json:"confirmed"`
	Recovered             int     `json:"recovered"`
	Population_percentage float64 `json:"population_percentage"`
}

//
func HandlerCoronaCase(w http.ResponseWriter, r *http.Request) {
	var err error

	urlArray := strings.Split(functions.GetURL(r), "/")
	//checks if required parameters are in place
	if len(urlArray) != 5 {
		fmt.Print(err.Error())
		return
	}
	//
	urlParameters, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	//
	country := urlArray[4]
	scope := ""
	//if "scope" parameter DOES exist
	if len(urlParameters) > 0 {
		if scopeParameter, ok := urlParameters["scope"]; ok {
			scope = scopeParameter[0]
		} else {
			//if error
			fmt.Print(err.Error())
			return
		}
	} else {
		//if "scope" parameter DOESN'T exist
		//scope is already empty if unchanged
	}
	//gets data from api
	var confirmedCases coronaCases
	var recoveredCases coronaCases
	err = getCoronaCases(&confirmedCases, &recoveredCases, country)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	//reformats and updates the date for the output
	var dataOutput outputCoronaCases
	updateOutputCoronaCases(&confirmedCases, &recoveredCases, &dataOutput, scope)
	//set header to json
	w.Header().Set("Content-Type", "application/json")
	//sends output
	err = json.NewEncoder(w).Encode(dataOutput)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

//
func getCoronaCases(confirmed *coronaCases, recovered *coronaCases, country string) error {
	var err error
	url := ""
	//gets confirmed cases data
	url = "https://covid-api.mmediagroup.fr/v1/history?country=" + country + "&status=Confirmed"
	confirmedOutput, err := requestRawData(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(confirmedOutput, &confirmed)
	if err != nil {
		return err
	}
	//gets recovered cases data
	url = "https://covid-api.mmediagroup.fr/v1/history?country=" + country + "&status=Recovered"
	recoveredOutput, err := requestRawData(url)
	if err != nil {
		return err
	}
	//outputs data to pointer
	err = json.Unmarshal(recoveredOutput, &recovered)
	return err
}

//
func updateOutputCoronaCases(confirmed *coronaCases, recovered *coronaCases, output *outputCoronaCases, scope string) {
	startDate := scope[:10]
	endDate := scope[11:]
	//country name
	output.Country = confirmed.All.Country
	//continent
	output.Continent = confirmed.All.Continent
	//scope
	output.Scope = scope
	//confirmed
	output.Confirmed = confirmed.All.Dates[endDate] - confirmed.All.Dates[startDate]
	//recovered
	output.Recovered = recovered.All.Dates[endDate] - recovered.All.Dates[startDate]
	//population_percentage
	output.Population_percentage = math.Round((float64(output.Confirmed)/float64(confirmed.All.Population)*100)*100) / 100
}
