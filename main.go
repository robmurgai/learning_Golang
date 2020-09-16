package main

import (
	"fmt"
)

// main is used to call all the expercise functions.
func main() {
	fmt.Printf("######### START ###########\n\n")

	//landing()

	var earth = world{radius: 6371.0} //Earth

	//Find the distance from London, England (51°30’N 0°08’W) to Paris, France (48°51’N 2°21’E).
	london := location{Name: "London, England", Lat: coordinate{51, 30, 0, 'N'}, Long: coordinate{0, 8, 0, 'W'}}
	paris := location{Name: "Paris, France", Lat: coordinate{48, 51, 0, 'N'}, Long: coordinate{2, 21, 0, 'E'}}
	fmt.Printf("London to Paris is: %.2f km as the bird flyes\n", earth.distance(london, paris))

	//Find the distance from your city to the capital of your country.
	// Cary is 35.7915° N, 78.7811° W
	// Washington, D.C./Coordinates 38.9072° N, 77.0369° W
	cary := location{Name: "Cary, NC", Lat: coordinate{d: 35.7915, h: 'N'}, Long: coordinate{d: 78.7811, h: 'W'}}
	washDC := location{Name: "Washington DC", Lat: coordinate{d: 38.9072, h: 'N'}, Long: coordinate{d: 77.0369, h: 'W'}}
	fmt.Printf("Distance between %v and %v is: %.2f km as the bird flyes", cary.Name, washDC.Name, earth.distance(cary, washDC))

	fmt.Printf("\n\n########## END ############\n")
}
