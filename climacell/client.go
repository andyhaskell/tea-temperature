package climacell

import (
	"encoding/json"
	"fmt"
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

	switch res.StatusCode {
	case 200:
		var weatherSamples []Weather
		if err := json.NewDecoder(res.Body).Decode(&weatherSamples); err != nil {
			return nil, err
		}
		return weatherSamples, nil
	case 400, 401, 403, 500:
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		// handle unexpected status codes
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// ErrorResponse represents a JSON error response from a ClimaCell API
// endpoint.
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	if err.ErrorCode == "" {
		return fmt.Sprintf("%d API error: %s", err.StatusCode, err.Message)
	}
	return fmt.Sprintf("%d (%s) API error: %s", err.StatusCode, err.ErrorCode, err.Message)
}
