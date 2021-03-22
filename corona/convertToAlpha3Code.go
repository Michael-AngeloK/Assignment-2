package corona

import (
	"encoding/json"
)

//
type countryAlpha []struct {
	AlphaCode string `json:"alpha3Code"`
}

// Outputs countryAlpha struct
func convertToAlphaCode(country string) (string, error) {
	var err error
	var inputData countryAlpha

	err = getAlphaCode(&inputData, country)
	if err != nil {
		return "", err
	}
	outputData := inputData[0].AlphaCode
	return outputData, err
}

//
func getAlphaCode(c *countryAlpha, country string) error {
	var err error
	url := "https://restcountries.eu/rest/v2/name/" + country + "?fields=alpha3Code"

	output, err := requestRawData(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(output, &c)
	return err
}
