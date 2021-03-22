package corona

import (
	"io/ioutil"
	"net/http"
	"time"
)

//requestRawData gets raw data
func requestRawData(url string) ([]byte, error) {
	var err error

	req, err := http.NewRequest(http.MethodGet, url, nil)
	client := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	output, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return output, err
}
