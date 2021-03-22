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
	http.HandleFunc("/corona/v1/country/", corona.HandlerCoronaCase)
	//corona Stingency handler
	http.HandleFunc("/corona/v1/policy/", corona.HandlerCoronaStringency)
	//diagnosis handler
	http.HandleFunc("/corona/v1/diag/", corona.HandlerDiagnosis)
	//notifications handler
	http.HandleFunc("/corona/v1/notifications/", corona.HandlerNotifications)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
