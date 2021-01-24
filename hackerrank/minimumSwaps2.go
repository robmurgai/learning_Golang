// I made two attempts at this
// my first attemt was MinimumSwaps() and my sencond attempt was MinimumSwaps1/
// While the first one worked, it did not take any of the limiations into considerations, like the unordered array had consecutive integers.
//

package hackerrank

import (
	"fmt"
)

// MinimumSwaps1 does the following
// for an array of n elements, 1 to n, ideal array is 1, 2, 3, ..., n.
// so arr[i] = i+1 in a sorted array
// and max swaps is len(arr) - 1
// go through the input array and see which elements are not in the correct order, these have to be swapped.
// for every incorrect arr[i] = j+1, swap arr[i] and arr[j]
// 		arr := []int32{4, 3, 1, 2}, max swaps 3
//		swap arr[0] and arr[3]: {2, 3, 1, 4}
//		swap arr[0] and arr[1]: {3, 2, 1, 4}
//		swap arr[0] and arr[2]: {1, 2, 3, 4}
//		return 3, you have rached max swaps.
//
//		arr := []int32{2, 3, 4, 1, 5}, max swaps is 4
//		swap arr[0] and arr[1]: {3, 2, 4, 1, 5}
//		swap arr[0] and arr[2]: {4, 2, 3, 1, 5}
//		swap arr[0] and arr [3]: {1, 2, 3, 4, 5}
//		check elements 1, 2, 3, 4, 5.
//		return 3
//
//		arr := []int32{1, 3, 5, 2, 4, 6, 7}, max swaps 6
//		swap arr[1] and arr[2]: {1, 5, 3, 2, 4, 6, 7}
//		swap arr[1] and arr[4]: {1, 4, 3, 2, 5, 6, 7}
//		swap arr[1] and arr[3]: {1, 2, 3, 4, 5, 6, 7}
//		check elements 3, 4, 5, 6, 7
//
func MinimumSwaps1(arr []int32) int32 {

	fmt.Printf("Input Array of %v elements: %v\n", len(arr), arr)

	numOfSwaps := int32(0)
	maxNumOfSwaps := int32(len(arr) - 1)

	fmt.Printf("numOfSwaps: %v\n", numOfSwaps)
	fmt.Printf("maxnumOfSwaps: %v\n", maxNumOfSwaps)

	for i, val := range arr {
		fmt.Printf("In the for loop.  Processing arr[%v]: %v\n", i, val)

		for val != int32(i+1) {
			arr[i] = arr[val-1]
			arr[val-1] = val
			numOfSwaps++
			val = arr[i]
			fmt.Printf("Swap: %v\n", numOfSwaps)
			fmt.Printf("Swapped Array of %v elements: %v\n", len(arr), arr)
		}

		if numOfSwaps == maxNumOfSwaps {
			fmt.Printf("Trying to break out of the for loop\n")
			break
		}

	}
	return numOfSwaps
}

// MinimumSwaps function below.
// Flow
// two cursors: left and right
//  move the left cursor from left to right, if
//      the next element is bigger than the current element.
//      and there is a next element
//      otherwise stop at current cursor.
// move the right cursor from right to left, if
//      the next element is smaller than the current element
//      and the left cursor is not next to the right cursor
//      otherwise stop.
// when both cursors stop or they have met,
// if left cursor value is bigger than right cursor value --> swap
// otherwise you are done.
func MinimumSwaps(arr []int32) int32 {

	fmt.Printf("Input Array of %v elements: %v\n", len(arr), arr)

	numOfSwaps := int32(0)
	fmt.Printf("numOfSwaps: %v\n", numOfSwaps)

	stillSwapping := true
	for stillSwapping {
		fmt.Printf("stillSwapping: %v\n", stillSwapping)

		// Move left cursor form left to right
		// Set stopLeftCursor to the 2nd last element of the array
		stopLeftCursor := len(arr) - 2
		tmpPrintString := ""

		for leftCursor := 0; leftCursor < stopLeftCursor; leftCursor++ {
			fmt.Printf("  Left Cursor: %v    Comaparing (%v,%v)\n", leftCursor, arr[leftCursor], arr[leftCursor+1])
			tmpPrintString = fmt.Sprintf("End of Array reached.")

			//stop at current element if the next element is smaller than the current element.
			if arr[leftCursor+1] < arr[leftCursor] {
				stopLeftCursor = leftCursor
				tmpPrintString = fmt.Sprintf("Left Cursor arr[%v]: %v is smaller than arr[%v]: %v.", leftCursor, arr[leftCursor], leftCursor+1, arr[leftCursor+1])
			}
		}
		fmt.Printf("    %v Stopping left cursor at arr[%v]: %v\n", tmpPrintString, stopLeftCursor, arr[stopLeftCursor])

		// move the right cursor from right to left
		stopRightCursor := stopLeftCursor + 1
		tmpPrintString = ""
		for rightCursor := len(arr) - 1; rightCursor > stopRightCursor; rightCursor-- {
			fmt.Printf("  Right Cursor: %v", rightCursor)
			fmt.Printf("    Comaparing (%v,%v)\n", arr[rightCursor-1], arr[rightCursor])
			//stop at the current element if
			//  the next element is bigger than the current element
			//  and
			//  the left cursor element is bigger than the current element
			if arr[rightCursor-1] > arr[rightCursor] {
				tmpPrintString = fmt.Sprintf("Right Cursor arr[%v]: %v is smaller than arr[%v]: %v, ", rightCursor-1, arr[rightCursor-1], rightCursor, arr[rightCursor])

				if arr[stopLeftCursor] > arr[rightCursor] {
					stopRightCursor = rightCursor
					tmpPrintString = fmt.Sprintf("Left Cursor arr[%v]: %v is bigger than the Right Cursor at arr[%v]: %v and \n    %v", stopLeftCursor, arr[stopLeftCursor], rightCursor, arr[rightCursor], tmpPrintString)

				} else {
					fmt.Printf("    %vbut\n    Left Cursor arr[%v]: %v is NOT bigger than Right Cursor arr[%v]: %v. Continue\n", tmpPrintString, stopLeftCursor, arr[stopLeftCursor], rightCursor, arr[rightCursor])
					tmpPrintString = ""

				}
			}
		}
		fmt.Printf("    %vstopping right cursor at arr[%v]: %v \n", tmpPrintString, stopRightCursor, arr[stopRightCursor])

		//if the value at stopleftCursor is greater than value at stopRightCursor, swap
		if arr[stopLeftCursor] > arr[stopRightCursor] {
			fmt.Printf("  Comparing (%v,%v) in Current Array of %v elements: %v\n", arr[stopLeftCursor], arr[stopRightCursor], len(arr), arr)
			fmt.Printf("  Swapping arr[%v]: %v and arr[%v]: %v\n", stopLeftCursor, arr[stopLeftCursor], stopRightCursor, arr[stopRightCursor])
			tmp := arr[stopRightCursor]
			arr[stopRightCursor] = arr[stopLeftCursor]
			arr[stopLeftCursor] = tmp
			numOfSwaps++
			fmt.Printf("\nSwapped Array of %v elements: %v\n", len(arr), arr)
			fmt.Printf("numOfSwaps: %v\n", numOfSwaps)
		} else {
			stillSwapping = false
		}
	}
	return numOfSwaps
}
