package main

import (
	//For Printing to the terminal
	"fmt"
)

func main() {

	fmt.Printf("######### START ###########\n\n")

	capture := dataCapture()
	fmt.Println("Debug Main(): Initalize capture - An indigoValue object")
	capture.prettyPrint()

	capture.add(3)
	capture.add(9)
	capture.add(3)
	capture.add(4)
	capture.add(6)
	// capture.add(0)
	// capture.add(500)
	// capture.add(1000)
	// capture.add(1001)
	// capture.add(1002)
	// capture.add(-1)

	fmt.Println("Debug Main(): Add 3, 9, 3, 4, 6")
	capture.prettyPrint()

	stats := capture.buildStats()
	fmt.Println("Debug main(): Initialize stats - An indigoStats Object")
	// fmt.Printf("%#v", stats)
	stats.prettyPrint()

	//stats.less(4) should return 2 (only two values 3, 3 are less than 4)
	fmt.Printf("stats.less(4): %v\n", stats.less(4))

	// stats.between(3, 6) # should return 4 (3, 3, 4 and 6 are between 3 and 6)
	fmt.Printf("stats.between(3, 6): %v\n", stats.between(3, 6))

	//# stats.less(4) should return 2 (6 and 9 are the only two values greater than 4)
	fmt.Printf("stats.greater(4): %v\n", stats.greater(4))

	fmt.Printf("\n\n########## END ############\n")
}
