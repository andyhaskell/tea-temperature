package climacell

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

//
// request query parameters
//

// LatLon represents a pair of latitude and longitude coordinates.
type LatLon struct{ Lat, Lon float64 }

// ForecastArgs represents query parameters that can be sent to the ClimaCell
// API when requesting weather data. We create query parameters with the
// QueryParams method.
type ForecastArgs struct {
	// If present, latitude and longitude coordinates we are requesting
	// forecast data for.
	LatLon *LatLon
	// If non-blank, ID for location we are requesting forecast data for.
	LocationID string
	// Unit system to return weather data in. Valid values are "si" and "us",
	// default is "si"
	UnitSystem string
	// Weather data fields we want returned in the response
	Fields []string
	// If nonzero, StartTime indicates the initial timestamp to request weather
	// data from.
	StartTime time.Time
	// If nonzero, EndTime indicates the ending timestamp to request weather
	// data to.
	EndTime time.Time
}

// QueryParams converts this ForecastArgs into a url.Values struct so it can be
// used as query parameters for weather requests. Each field on the
// ForecastArgs struct is only converted to a query parameter if that field is
// not its zero value.
func (args ForecastArgs) QueryParams() url.Values {
	q := make(url.Values)

	if args.LatLon != nil {
		q.Add("lat", strconv.FormatFloat(args.LatLon.Lat, 'f', -1, 64))
		q.Add("lon", strconv.FormatFloat(args.LatLon.Lon, 'f', -1, 64))
	}

	if args.LocationID != "" {
		q.Add("location_id", args.LocationID)
	}
	if args.UnitSystem != "" {
		q.Add("unit_system", args.UnitSystem)
	}

	if len(args.Fields) > 0 {
		q.Add("fields", strings.Join(args.Fields, ","))
	}

	if !args.StartTime.IsZero() {
		q.Add("start_time", args.StartTime.Format(time.RFC3339))
	}
	if !args.EndTime.IsZero() {
		q.Add("end_time", args.EndTime.Format(time.RFC3339))
	}

	return q
}

//
// response deserialization
//

// FloatValue is a field on a weather sample returned from the ClimaCell API
// whose type is a floating-point number. If Value is nil, that means that for
// the field this FloatValue is for, data is not available.
type FloatValue struct {
	Value *float64 `json:"value"`
	Units string   `json:"units"`
}

// NonNullableTimeValue is a field on a Weather for timestamps. This type is
// used for fields that should always be present on a weather sample, such as
// observation timestamps.
type NonNullableTimeValue struct {
	Value time.Time `json:"value"`
}

// Weather represents a weather sample from the ClimaCell API. Note that in the
// full API, there are many more fields that can be present, but for this
// tutorial we are only deserializing temperature.
type Weather struct {
	Lat             float64              `json:"lat"`
	Lon             float64              `json:"lon"`
	Temp            *FloatValue          `json:"temp"`
	ObservationTime NonNullableTimeValue `json:"observation_time"`
}
