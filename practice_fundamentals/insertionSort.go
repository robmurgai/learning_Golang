package practicefundamentals

import "fmt"

// Insertion sort is a step more complex but a bit more useful than bubble sort and is occasionally useful. The worst case
// scenario for it is similar to bubble sort's but its best case makes it suited for times when you're pretty sure a list ////
// almost sorted or likely already sorted.

// We're going to start at the beginning of the list and assume we have a sorted list of length 1 where the first element is
// only sorted element. We're then going to grab the second element, and insert it into the correct spot in our sorted list, //
// either the 0 index or the 1 index, depending if it's smaller or larger than our first element. We now have a sorted list of /// length 2. We then continue on down the line, inserting elements in our sorted side of the list as the unsorted side dwindles.

// What's the Big O? There's an inner loop that goes over your sorted list to find the correct place to insert your item, and
// an outer loop to go over all the numbers. Two loops means O(n²). However since if your list is sorted or nearly so, it can /
// be O(n) in a best case scenario and thus well adapted to that scenario.

func insertionSort(myList []int) {

	bigO := 0

	newList := make([]int, 0, cap(myList))
	fmt.Printf("newList: %v\n", newList)

	newList = append(newList, myList[0])
	fmt.Printf("newList: %v\n", newList)

	for _, valueFromOldList := range myList[1:] {
		bigO++
		for index, valueInNewList := range newList {
			bigO++
			if valueFromOldList < valueInNewList {
				fmt.Printf("valueFromOldList: %v valueInNewList: %v\n", valueFromOldList, valueInNewList)
				// Insert valueFromOldList to the left of index
				// newlist[:index] + valueFromOldList + newList[index+1:]
				newList = insertIntoList(newList, valueFromOldList, index)
				break
			}
		}
	}

	for index, value := range newList {
		myList[index] = value
	}

	fmt.Printf("n: %v, BigO(n): %v\n", len(myList), bigO)

}

func insertIntoList(theList []int, value int, position int) []int {

	endList := make([]int, 0, cap(theList))

	endList = append(endList, theList[:position]...)
	endList = append(endList, value)
	endList = append(endList, theList[position:]...)

	return endList

}
