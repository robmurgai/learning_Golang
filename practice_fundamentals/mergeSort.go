package practicefundamentals

import "fmt"

// The basic gist of merge sort is that you're going to take your big list, and first divide down in two half size lists and
// recursively call merge sort on those smaller list, which in turn will do the same. The base case is when you have a list of
// one, at which point you will return that sorted list of one. On the way up the recursive calls, you will merge those sorted
// lists together (preferably by another merge function you'll write) that walks through both lists simultaneously and inserts
// the smaller value first, effectively creating a bigger sorted list.

// [1, 5, 6] sublist 1
// [2, 7, 8] sublist 2

// -> compare 1 and 2, take 1 and put it in new list
// -> compare 5 and 2, take 2 and put it in new list
// -> compare 5 and 7, take 5 and put it in new list
// -> compare 6 and 7, take 6 and put it in new list
// -> list one has no more elements
//    add the rest of list two in order (7 and 8)

// MergeSort's Big O is O(n log n). Weird, right? We obviously have to compare everything once, but we don't have to compare
// everything to everything like we do with bubble sort. Rather we just to have to compare to their local lists as we sort.
// Not too bad.

// MergeSort's space complexity is a bit worse than the previous algorithms at O(n) since we have to create new lists as we
// go. It's not awful but it nonetheless a consideration.
func mergeSort(myList []int) {

	fmt.Printf("mergeSort this list: %v\n", myList)

	theLength := len(myList)
	endList := make([]int, theLength)

	if theLength > 1 {
		mergeSort(myList[:theLength/2])
		mergeSort(myList[(theLength / 2):])

		endList = mergeLists(myList[:theLength/2], myList[(theLength/2):])
		fmt.Printf("endList: %v\n", endList)

		for index, value := range endList {
			myList[index] = value
		}
	}
	fmt.Printf("myList: %v\n", myList)
}

func mergeLists(leftList []int, rightList []int) []int {

	fmt.Printf("\nMerge these two lists: %v & %v\n", leftList, rightList)

	mergedList := make([]int, 0, len(leftList)+len(rightList))
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(leftList) {
		bigO++
		leftValue := leftList[leftIndex]
		fmt.Printf("leftValue: %v & ", leftValue)

		for rightIndex < len(rightList) {
			bigO++
			rightValue := rightList[rightIndex]
			fmt.Printf("rightValue: %v\n", rightValue)

			if leftValue < rightValue {
				mergedList = append(mergedList, leftValue)
				leftIndex++
				break
			} else {
				mergedList = append(mergedList, rightValue)
				rightIndex++
			}
			fmt.Printf("mergedList: %v\n", mergedList)
		}
		if rightIndex == len(rightList) {
			fmt.Printf("Append rest of the LeftList: %v\n", leftList[leftIndex:])
			mergedList = append(mergedList, leftList[leftIndex:]...)
			fmt.Printf("mergedList: %v\n", mergedList)
			leftIndex = len(leftList)
		}
	}
	if rightIndex != len(rightList) {
		fmt.Printf("Append rest of the rightList: %v\n", rightList[rightIndex:])
		mergedList = append(mergedList, rightList[rightIndex:]...)
		fmt.Printf("mergedList: %v\n", mergedList)
		rightIndex = len(rightList)
	}

	return mergedList

}
