// Create an HTTP Server to understand how it works.

package main

import (
	"fmt"
	"log"
	"os"

	hr "./hackerrank"
)

var debugLog *log.Logger

func init() {

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	var n int32
	var queries [][]int32

	// Problem 1
	n = 5

	queries = [][]int32{[]int32{1, 2, 100}, []int32{2, 5, 100}, []int32{3, 4, 100}}

	fmt.Printf("Calling ArrayManipulation with querries: %v\n\n", queries)

	res := hr.ArrayManipulation(n, queries)

	fmt.Printf("Max Value: %d\n", res)

	fmt.Printf("\n########## Ending Exercise ##########\n")
}
