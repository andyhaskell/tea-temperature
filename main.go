package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://api.climacell.co/v3/weather/forecast/hourly?lat=42.3826&lon=-71.1460&fields=temp",
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", os.Getenv("CLIMACELL_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))
}
