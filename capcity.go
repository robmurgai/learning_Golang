/**
 ** Experiment: capacity.go
 ** Write a program that uses a loop to continuously append an element to a slice. Print out the capacity of the slice whenever it changes. Does append
 ** always double the capacity when the underlying array runs out of room?
 **/

package main

import "fmt"

func capacity() {

	planetSplice := []string{"Mercury", "Venus", "Earth"}

	fmt.Printf("planetSplice: %v of lenght %v and capacity %v\n\n", planetSplice, len(planetSplice), cap(planetSplice))

	for i, planet := range planetSplice {

		planetSplice = append(planetSplice, string(i)+planet)
		fmt.Printf("planetSplice[%v]: %v of lenght %v and capacity %v\n", i, planetSplice, len(planetSplice), cap(planetSplice))

		if i == 100 {
			break
		}
	}

	fmt.Println()

	for i := 0; i < len(planetSplice); i++ {
		planetSplice = append(planetSplice, string(i)+planetSplice[i])
		fmt.Printf("planetSplice[%v]: %v of lenght %v and capacity %v\n", i, planetSplice, len(planetSplice), cap(planetSplice))

		if i == 10 {
			break
		}

	}
}
