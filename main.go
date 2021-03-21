package main

import (
	"log"
	"main/corona"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	corona.StartTime = time.Now()
	//corona Cases handler
	http.HandleFunc("/corona/v1/country/", corona.HandlerCoronaCases)
	//corona Stingency handler
	http.HandleFunc("/corona/v1/policy/", corona.HandlerCoronaStingency)
	//diagnosis handler
	http.HandleFunc("/corona/v1/diag/", corona.HandlerDiagnosis)
	//notifications handler
	http.HandleFunc("/corona/v1/notifications/", corona.HandlerNotifications)
	//ends program if it can't open port
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
