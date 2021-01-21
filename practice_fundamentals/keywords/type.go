package keywords

import "fmt"

func typeKeyword() {

	// Create a type farenheight that stores temperature as a float64
	type farenheight float64

	var temperature1 farenheight
	fmt.Println("var temperature1 farenheight")
	fmt.Printf("temperature1: %v of type: %[1]T\n\n", temperature1)

	temperature2 := 3.54
	fmt.Println("temperature2 := 3.54")
	fmt.Printf("temperature2: %v of type: %[1]T\n\n", temperature2)

	temperature3 := farenheight(3.54)
	fmt.Println("temperature3 := farenheight(3.54)")
	fmt.Printf("temperature3: %v of type: %[1]T\n\n", temperature3)
}
