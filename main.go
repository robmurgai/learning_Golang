package main

import (
	"encoding/json"
	"fmt"
)

// main is used to call all the expercise functions.
func main() {
	fmt.Printf("######### START ###########\n\n")

	elysium := location{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}

	//fmt.Printf("%v\n\n", elysium)

	bytes, err := json.MarshalIndent(elysium, "", "  ")
	exitOnError(err)
	fmt.Println(string(bytes))

	fmt.Printf("\n\n########## END ############\n")
}
