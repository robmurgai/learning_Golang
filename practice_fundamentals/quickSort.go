package practicefundamentals

// Quicksort is one of the most useful and powerful sorting algorithms out there, and it's not terribly difficult to
// conceptualize (compared to some algorithms we're not talking about anyway.)
// JavaScript doesn't mergesort for Array.prototype.sort. In those other cases, it's usually some variant on quicksort.

// It's another divide-and-conquer, recursive algorithm but it takes a slightly different approach. The basic gist is that
// you take the last element in the list and call that the pivot. Everything that's smaller than the pivot gets put into
// a "left" list and everything that's greater get's put in a "right" list.
//
// You then call quick sort on the left and right lists independently (hence the recursion.) After those two sorts come back,
// you concatenate the sorted left list, the pivot, and then the right list (in that order.) The base case is when you have
// a list of length 1 or 0, where you just return the list given to you.

// [4,9,3,5] list
// -> 5 is made the pivot since it's the last in the array
// -> divide list into two lists, [4,3] and [9]
// -> call quicksort on those two lists

// [4, 3]
// -> 3 is pivot
// -> call quicksort on [] and [4]
// -> those both return as is as they are the base case of length 0 or 1
// -> concat [], 3, and [4]
// -> return [3,4]

// [9]
// -> returns as this it is a base case of length 1

// (back into the original function call)
// -> call concat on [3,4], 5, and [9]
// -> return [3,4,5,9]

// Another Big O of O(n log n) but takes up less memory than mergesort so it is often favored. However it does really poorly
// if you pass it a sorted list. Think about it. It would always have a pivot of the biggest number which defeats the effectiveness of the divide-and-conquer approach as one side will always contain all the elements. Hence not good for lists
// you expect may already be sorted. There are some tricks to employ to get around that like checking the beginning, middle,
// and end numbers and swapping them to try to get the best pivot, but that's outside our scope today. There are a lot of
// subtle variants on quicksort.

func QuickSort(myList []int) {

	if len(myList) > 1 {
		theLength := len(myList)
		pivot := myList[theLength-1]
		leftList := make([]int, 0, theLength)
		rightList := make([]int, 0, theLength)

		for _, value := range myList[:theLength-1] {
			if value < pivot {
				leftList = append(leftList, value)
			} else {
				rightList = append(rightList, value)
			}
		}

		QuickSort(leftList)
		QuickSort(rightList)

		leftList = append(leftList, pivot)
		leftList = append(leftList, rightList...)
		for index, value := range leftList {
			myList[index] = value
		}
	}

}
