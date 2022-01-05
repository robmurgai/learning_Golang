package practicefundamentals

import "fmt"

// Our first sort! This is often the easiest to conceptualize and a natural way for the brain to think about sorting so it's
// typical to do bubble sort first. It's also amongst the least efficient in terms of worst case scenario.

// In bubble sort, we're going to loop through the array and compare each index with the index next to it. If the those two
// numbers are out of order (the lesser index's value is greater than the greater index's value) we swap those two numbers'
// places in the array. We keep looping over that array until everything is in place and nothing was swapped during the last
// iteration.

// What's the Big O on this? Well, there's an inner loop to check to see if indexes need to be swapped, and an outer loop
// that's just checking to see if anything was swapped. That would be make it O(n²). Not efficient, but a great learning tool.
// You'll never use bubble sort for anything serious.

/**
To run BubbleSort, update main.go

package main

import (
	pf "github.com/robmurgai/learning_Golang/practice_fundamentals"
)

func main() {
	myList := []int{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	pf.BubbleSort(myList)
}

**/

// BubbleSort() sorts an array and prints out the sorted array.
// Expected input {10, 5, 3, 8, 2, 6, 4, 7, 9, 1} to equal {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func BubbleSort(myList []int) {

	keepGoing := true
	outerLoop := 1
	innerLoop := 1

	for keepGoing {
		keepGoing = false

		fmt.Printf("Outer Loop (%v): %v\n", outerLoop, myList)
		outerLoop++

		innerLoop = 1
		// Range from index 0 to 'last but 1' because we are comparing myList[index] with myList[index+1]
		for index, value := range myList[:len(myList)-1] {
			fmt.Printf("  %v: %v\n", innerLoop, myList)
			innerLoop++
			if myList[index+1] < value {
				myList[index] = myList[index+1]
				myList[index+1] = value
				keepGoing = true
			}
		}

	}

	outerLoop--

	//fmt.Printf("OuterLoop: %v and innerLoop: %v\n", outerLoop, innerLoop)
	fmt.Printf("BubbleSort for n:%v is BigO(n^2): %v\n", len(myList), outerLoop*innerLoop)

}
