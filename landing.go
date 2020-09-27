package main

import (
	"fmt"
)

// landing declares a location for each roq in the following table and prints out each of the locations in decimal degrees.
//
// Rover or lander			Landing site				Latitude			Longitude
// Spirit					Columbia Memorial Station	14°34’6.2” S		175°28’21.5” E
// Opportunity				Challenger Memorial Station	1°56’46.3” S		354°28’24.2” E
// Curiosity				Bradbury Landing			4°35’22.2” S		137°26’30.1” E
// InSight					Elysium Planitia			4°30’0.0” N			135°54’0” E
func landing() {

	// Mars the World
	var mars = world{radius: 3389.5} //Mars

	// Spirit Columbia Memorial Station	14°34’6.2” S	175°28’21.5”
	Spirit := location{Name: "Columbia Memorial Station", Lat: coordinate{14, 34, 6.2, 'S'}, Long: coordinate{175, 28, 21.5, 'E'}}
	fmt.Printf("Spirit landed at %v\n", Spirit.locationInDecimal())

	// Opportunity	Challenger Memorial Station	1°56’46.3” S	354°28’24.2” E
	Oppertunity := location{Name: "Challenger Memorial Station", Lat: coordinate{1, 56, 46.3, 'S'}, Long: coordinate{354, 28, 24.2, 'E'}}
	fmt.Printf("Oppertunity landed at %v\n", Oppertunity.locationInDecimal())

	// Curiosity	Bradbury Landing	4°35’22.2” S	137°26’30.1” E
	Curiosity := location{Name: "Bradbury Landing", Lat: coordinate{4, 35, 22.2, 'S'}, Long: coordinate{137, 26, 30.1, 'E'}}
	fmt.Printf("Curiosity landed at %v\n", Curiosity.locationInDecimal())

	// InSight	Elysium Planitia	4°30’0.0” N	135°54’0” E
	InSight := location{Name: "Elysium Planitia", Lat: coordinate{4, 30, 0.0, 'N'}, Long: coordinate{135, 54, 0, 'E'}}
	fmt.Printf("InSight landed at %v\n", InSight.locationInDecimal())

	fourMarsSites := []location{Spirit, Oppertunity, Curiosity, InSight}

	for i, foo := range fourMarsSites {

		for _, bar := range fourMarsSites[i+1:] {
			fmt.Printf("Distance between %v and %v is: %.2f\n", foo.Name, bar.Name, mars.distance(foo, bar))
		}
	}
}
