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

// NB!: NEEDS DATE VERIFICATION, ALPHA3CODE, COUNTRY, SCOPE = -1, ERRORS,

//
func HandlerCoronaStringency(w http.ResponseWriter, r *http.Request) {
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
	var country []string
	country = append(country, urlArray[4])
	//convert to alpha code
	country[0], err = convertToAlphaCode(country[0])
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
	var stringencyData coronaStringency
	err = getCoronaStringency(&stringencyData, scope)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	//reformats and updates the date for the output
	var dataOutput outputCoronaStringency
	updateOutputCoronaStringency(&stringencyData, &dataOutput, country[0], scope)
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
func getCoronaStringency(corona *coronaStringency, scope string) error {
	var err error
	url := ""
	//dates
	startDate := scope[:10]
	endDate := scope[11:]

	//gets confirmed cases data
	url = "https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/date-range/" + startDate + "/" + endDate
	confirmedOutput, err := requestRawData(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(confirmedOutput, &corona)
	return err
}

//
func updateOutputCoronaStringency(stringency *coronaStringency, output *outputCoronaStringency, alphaCode string, scope string) {
	//dates
	startDate := scope[:10]
	endDate := scope[11:]
	//country name
	output.Country = alphaCode
	//scope
	output.Scope = scope
	//confirmed
	output.Stringency = math.Round(stringency.Data[endDate][alphaCode].Stringency*100) / 100
	//recovered
	output.Trend = math.Round((stringency.Data[endDate][alphaCode].Stringency-stringency.Data[startDate][alphaCode].Stringency)*100) / 100
}
