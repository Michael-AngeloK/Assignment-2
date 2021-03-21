package corona

import (
	"encoding/json"
)

//
type countryAlpha []struct {
	Alpha3Code string `json:"alpha3Code"`
}

//
func handlerConvertToAlpha(country string) (string, error) {
	var err error
	var inpData countryAlpha

	err = getCountryAlpha(&inpData, country)
	if err != nil {
		return "", err
	}
	outData := inpData[0].Alpha3Code
	return outData, err
}

//
func getCountryAlpha(c *countryAlpha, country string) error {
	var err error
	url := "https://restcountries.eu/rest/v2/name/" + country + "?fields=alpha3Code"

	output, err := requestRawData(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(output, &c)
	return err
}
