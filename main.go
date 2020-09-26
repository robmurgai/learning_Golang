package main

import (
	"fmt"
)

// main is used to call all the expercise functions.
func main() {
	fmt.Printf("######### START ###########\n\n")

	//rover
	var rover rover

	//set up the planet rover is on
	var earth = world{radius: 6371.0} //Earth
	rover.gps.world = earth

	//Create two locations on the world.
	cary := location{Name: "Cary, NC", Lat: coordinate{d: 35.7915, h: 'N'}, Long: coordinate{d: 78.7811, h: 'W'}}
	washDC := location{Name: "Washington DC", Lat: coordinate{d: 38.9072, h: 'N'}, Long: coordinate{d: 77.0369, h: 'W'}}
	fmt.Printf("Distance between %v and %v is: %.2f km as the bird flyes\n", cary.Name, washDC.Name, earth.distance(cary, washDC))

	rover.current = cary
	rover.desitination = washDC

	fmt.Printf(rover.message())

	fmt.Printf("\n\n########## END ############\n")
}
