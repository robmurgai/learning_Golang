package main

import (
	"fmt"

	pf "github.com/robmurgai/learning_Golang/practiceFundamentals"
)

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	//Unsorted List
	myList := []int{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	fmt.Printf("myList before sort: %v\n", myList)
	pf.QuickSort(myList)
	fmt.Printf("myList after sort: %v\n", myList)
	fmt.Printf("n: %v, BigO(n): %v\n", len(myList), bigO)

	//Partially Sorted List
	//Sorted List

	fmt.Printf("\n########## Ending Exercise ##########\n")
}
