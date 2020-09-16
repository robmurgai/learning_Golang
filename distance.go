package main

import "math"

type world struct {
	radius float64
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance determines the distance between each pair of landing sites using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat.decimal()))
	s2, c2 := math.Sincos(rad(p2.Lat.decimal()))
	clong := math.Cos(rad(p1.Long.decimal() - p2.Long.decimal()))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
