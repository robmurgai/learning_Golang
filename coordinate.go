package main

import "fmt"

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

//     "decimal": 135.9,
//     "dms": "135°54'0.0\" E",
//     "degrees": 135,
//     "minutes": 54,
//     "seconds": 0,
//     "hemisphere": "E"
//String returns a string with the following DMS: "135°54'0.0 "E" and Decimal Values
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f \"%c", c.d, c.m, c.s, c.h)
}
