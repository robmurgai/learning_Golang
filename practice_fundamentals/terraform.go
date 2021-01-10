package practicefundamentals

import "fmt"

//Planets exported type
type Planets []string

/**
 **
 ** Experiment: terraform.go
 ** Write a program to terraform a slice of strings by prepending each planet with "New ". Use your program to terraform Mars, Uranus, and Neptune.
 **
 ** Your first iteration can use a terraform function, but your final implementation should introduce a Planets type with a terraform method, similar
 ** to sort.StringSlice.
 **
 ** Methods are defined as func (variable_name new_type) name_of_the_function(Parameters) result(s) {}

 ** In the main fucnation, you would call it this way
 ** func main() {
 ** 	var planets Planets = []string{"mercury", "earth", "mars", "venus", "saturn", "jupiter", "uranus", "neptue"}
 ** 	planets.terrafrom()
 ** 	fmt.Printf("\n\n########## END ############\n")
 ** }
 **
 **/

func (planets Planets) terrafrom() {

	fmt.Printf("Before: %v\n", planets)
	fmt.Println("Terraformaing the Planets:- ")
	for i := range planets {
		planets[i] = "New " + planets[i]
	}

	fmt.Printf("After: %v\n", planets)

}

func terrafrom1stIteration() {

	planets := []string{"mercury", "earth", "mars", "venus", "saturn", "jupiter", "uranus", "neptue"}
	fmt.Printf("Before: %v\n", planets)
	fmt.Println("Terraformaing the Planets:- ")
	for i := range planets {
		planets[i] = "New " + planets[i]
	}

	fmt.Printf("After: %v\n", planets)

}
