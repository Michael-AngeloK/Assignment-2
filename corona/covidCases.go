package corona

import (
	"encoding/json"
	"fmt"
	"main/functions"
	"net/http"
	"net/url"
	"strings"
)

//
type coronaCases struct {
	All struct {
		Country             string         `json:"country"`
		Population          int            `json:"population"`
		Sq_km_area          int            `json:"sq_km_area"`
		Life_expectancy     string         `json:"life_expectancy"`
		Elevation_in_meters int            `json:"elevation_in_meters"`
		Continent           string         `json:"continent"`
		Abbrevation         string         `json:"abbreviation"`
		Location            string         `json:"location"`
		Iso                 int            `json:"iso"`
		Capital_city        string         `json:"capital_city"`
		Dates               map[string]int `json:"dates"`
	} `json:"All"`
}

//
type outputCoronaCases struct {
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

	// Norway, scope=2020-12-01-2021-01-31
	urlParameters, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

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
	var confirmedCases coronaCases
	var recoveredCases coronaCases
	err = getCoronaCases(&confirmedCases, &recoveredCases, country)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	//var dataOutput coronaCases
	//err = getCoronaCases(&dataOutput, date)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(confirmedCases)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

//
func getCoronaCases(c *coronaCases, r *coronaCases, country string) error {
	var err error
	url := ""
	//startDate := date[:10]
	//endDate := date[11:]

	url = "https://covid-api.mmediagroup.fr/v1/history?country=" + country + "&status=Confirmed"
	confirmedOutput, err := requestRawData(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(confirmedOutput, &c)
	if err != nil {
		return err
	}
	url = "https://covid-api.mmediagroup.fr/v1/history?country=" + country + "&status=Recovered"
	recoveredOutput, err := requestRawData(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(recoveredOutput, &r)
	return err
}
