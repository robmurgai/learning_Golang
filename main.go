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

//To work with universe.go
// //create a NewUniverse and Show it.
// newUniverse := NewUniverse()
// tempUniverse := NewUniverse()
// //newUniverse.Show()
// //Seed the Universe and Show it.
// newUniverse.Seed()
// newUniverse.Show()
// //Clear the Screen on MAC
// //fmt.Println("\033[H")
// //Step the Universe and Show it.
// Step(newUniverse, tempUniverse)
// newUniverse = tempUniverse
// //fmt.Printf("\033[H")
// newUniverse.Show()

// To work with Words.go
// textArgument := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead
// the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he
// ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra.
// Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except
// the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
// words(textArgument)
