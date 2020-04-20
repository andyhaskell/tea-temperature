package main

import (
	"log"
	"os"
	"time"

	"github.com/andyhaskell/tea-temperature/climacell"
)

func main() {
	c := climacell.New(os.Getenv("CLIMACELL_API_KEY"))
	weatherSamples, err := c.HourlyForecast(climacell.ForecastArgs{
		LatLon:     &climacell.LatLon{Lat: 42.3826, Lon: -71.146},
		UnitSystem: "us",
		Fields:     []string{"temp"},
		StartTime:  time.Now(),
		EndTime:    time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		log.Fatalf("error getting forecast data: %v", err)
	}

	var tempAtFive *climacell.FloatValue
	for i, w := range weatherSamples {
		if w.ObservationTime.Value.Hour() == 21 {
			tempAtFive = weatherSamples[i].Temp
			break
		}
	}

	if tempAtFive == nil || tempAtFive.Value == nil {
		log.Printf("No data on the temperature at 5, let's wing it! 🌺\n")
	} else if t := *tempAtFive.Value; t < 60 {
		log.Printf("It'll be %f out. Better make some hot tea! 🌺🍵\n", t)
	} else {
		log.Printf("It'll be %f out. Iced tea it is! 🌺🍹\n", t)
	}
}
