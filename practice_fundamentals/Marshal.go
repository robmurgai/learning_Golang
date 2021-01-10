package practicefundamentals

import (
	"encoding/json"
)

// MarshalJSON outputs coordinates in JSON format, expanding on work done for the preceding quick check.
//The JSON output should provide each coordinate in decimal degrees (DD) as well as the degrees, minutes, seconds
// format:

// {
//     "decimal": 135.9,
//     "dms": "135°54'0.0\" E",
//     "degrees": 135,
//     "minutes": 54,
//     "seconds": 0,
//     "hemisphere": "E"
// }
// This can be achieved without modifying the coordinate structure by satisfying the json.Marshaler interface to
// customize the JSON. The MarshalJSON method you write may make use of json.Marshal.
// To calculate decimal degrees, you’ll need the decimal method introduced in lesson 22.
// String formats a DMS coordinate.
func (c coordinate) MarshalJSON() ([]byte, error) {

	//locationJSON created for this Marshalling function only
	// type cJSON struct {
	// 	DD  float64 `json:"decimal"`
	// 	DMS string  `json:"dms"`
	// 	D   float64 `json:"degrees"`
	// 	M   float64 `json:"minutes"`
	// 	S   float64 `json:"seconds"`
	// 	H   string  `json:"hemisphere"`
	// }
	type cJSON struct {
		DD  float64
		DMS string
		D   float64
		M   float64
		S   float64
		H   string
	}
	jsonByte := cJSON{DD: c.decimal(), DMS: c.String(), D: c.d, M: c.m, S: c.s, H: string(c.h)}

	//json.Marshall takes a struct and returns it as a byte array in JSON format.
	return json.Marshal(jsonByte)
}
