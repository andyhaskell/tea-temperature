package climacell

import (
	"time"
)

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
