package climacell

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.climacell.co",
	Path:   "/v3/",
}

// Client is a client for sending requests to the ClimaCell API.
type Client struct {
	c      *http.Client
	apiKey string
}

// New creates a new Go client for the ClimaCell API.
func New(apiKey string) *Client {
	c := &http.Client{Timeout: 30 * time.Second}

	return &Client{
		c:      c,
		apiKey: apiKey,
	}
}

// HourlyForecast gets an hourly forecast for a location.
func (c *Client) HourlyForecast(args ForecastArgs) ([]Weather, error) {
	// set up a request to the hourly forecast endpoint
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "weather/forecast/hourly"})
	req, err := http.NewRequest("GET", endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	// add URL headers, query params, then send the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", c.apiKey)
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	// deserialize the response and return our weather data
	defer res.Body.Close()
	var weatherSamples []Weather
	if err := json.NewDecoder(res.Body).Decode(&weatherSamples); err != nil {
		return nil, err
	}
	return weatherSamples, nil

	// [TODO] handle error responses
}
