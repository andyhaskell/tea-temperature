package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/andyhaskell/tea-temperature/climacell"
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

	var weatherSamples []climacell.Weather
	d := json.NewDecoder(res.Body)
	if err := d.Decode(&weatherSamples); err != nil {
		log.Fatalf("error deserializing weather data")
	}

	for _, w := range weatherSamples {
		if w.Temp != nil && w.Temp.Value != nil {
			log.Printf("The temperature at %s is %f degrees %s\n",
				w.ObservationTime.Value, *w.Temp.Value, w.Temp.Units)
		} else {
			log.Printf("No temperature data available at %s\n",
				w.ObservationTime.Value)
		}
	}
}
