package corona

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
type countryAlpha []struct {
	AlphaCode string `json:"alpha3Code"`
}

//
type coronaStringency struct {
	Data map[string](map[string]Stringency) `json:"data"`
}
type Stringency struct {
	Stringency_actual float64 `json:"stringency_actual"`
	Stringency        float64 `json:"stringency"`
}

//
type outputCoronaStringency struct {
	Country    string  `json:"country"`
	Scope      string  `json:"scope"`
	Stringency float64 `json:"stringency"`
	Trend      float64 `json:"trend"`
}

//diagnosis struct
type diagnosis struct {
	Mmediagroupapi  int    `json:"mmediagroupapi"`
	Covidtrackerapi int    `json:"covidtrackerapi"`
	Registered      int    `json:"registered"`
	Version         string `json:"version"`
	Uptime          string `json:"uptime"`
}
