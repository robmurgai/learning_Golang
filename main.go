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

	//arr := []int32{4, 3, 1, 2}
	//arr := []int32{2, 3, 4, 1, 5}
	//arr := []int32{1, 3, 5, 2, 4, 6, 7}
	arr := []int32{7, 1, 2, 3, 4, 5, 6}

	res := hr.MinimumSwaps1(arr)

	fmt.Printf("%d\n", res)

	fmt.Printf("\n########## Ending Exercise ##########\n")
}
