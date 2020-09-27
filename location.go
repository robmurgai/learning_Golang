package main

import (
	"fmt"
	"strconv"
)

// location in Name of the place, Latitude and Longitude both as coordinates
type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

type locationDecimal struct {
	Name string  `json:"name"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func (l location) locationInDecimal() string {
	return l.Name + " and " + strconv.FormatFloat(l.Lat.decimal(), 'f', -1, 64) + ", " + strconv.FormatFloat(l.Long.decimal(), 'f', -1, 64)
}

// String formats a location with latitude, longitude.
func (l location) String() string {
	returnString := fmt.Sprintf("%v is located at Latitude: %+v & Longitude: %+v", l.Name, l.Lat, l.Long)
	//returnString = fmt.Sprintf("%+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v", l.Lat.d, l.Lat.h, l.Lat.m, l.Lat.s, l.Lat.d, l.Lat.h, l.Lat.m, l.Lat.s)
	return returnString
}
