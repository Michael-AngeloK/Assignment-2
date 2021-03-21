package functions

import (
	"net/http"
)

//
func GetURL(r *http.Request) string {
	return r.URL.String()
}
