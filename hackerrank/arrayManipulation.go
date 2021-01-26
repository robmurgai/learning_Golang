package hackerrank

import "fmt"

// ArrayManipulation Second Attempt
// n is the size of the array and m is the number of operations.
// Each of the next m lines contains three space-separated integers a, b and k, the left index, right index and summand.
// first for loop is O(m), 2nd for loop is O(n) so O(m+n)
func ArrayManipulation(n int32, queries [][]int32) int64 {

	//Declare and Initialize
	//Lets say I am hiking and the array is my hike route
	//and every value in the array is the change in altitude.
	//so a value of 0 means no change in altitude, a value of 5 means, we increased the altitude by 5, etc
	//so the array is tracking the change in altitude and not the actual altitude.
	//in attempt 1, I was sorta tracking the altitude, in attmept 2 tracking the change in altitude
	//we are trying to see the max altitue the hike reaches
	myRoute := make([]int64, n)

	//The maximum altitude of the route starts at 0

	fmt.Println(myRoute)

	for _, value := range queries {
		// value: {a, b, k}
		// lets say k is the altitude increase and it increases at point 'a' in the route and decrease at point 'b+1' in the route
		// so myRoute[a-1] +=k and myRoute[b] -=k
		a := value[0]
		b := value[1]
		k := value[2]
		myRoute[a-1] += int64(k)
		if b < n {
			myRoute[b] -= int64(k)
		}

		fmt.Printf("%v: %v\n", value, myRoute)
	}

	maxAltitude := int64(0)
	altitude := int64(0)

	for _, value := range myRoute {
		altitude += value
		if maxAltitude < altitude {
			maxAltitude = altitude
		}
	}

	return maxAltitude

}

// ArrayManipulation1 First Attempt
// n is the size of the array and m is the number of operations.
// Each of the next m lines contains three space-separated integers a, b and k, the left index, right index and summand.
// first for loop is O(m), 2nd for loop is max O(n) so O(m*n)
func ArrayManipulation1(n int32, queries [][]int32) int64 {

	//Declare and Initialize
	myArray := make([]int64, n)

	maxArrayValue := myArray[0]

	fmt.Println(myArray)
	for _, value := range queries {
		// value: {a, b, k}
		// Add the values of k between the indices a and b inclusive
		tmpArray := myArray[value[0]-1 : value[1]]
		for tmpIndex := range tmpArray {
			tmpArray[tmpIndex] += int64(value[2])
			if maxArrayValue < tmpArray[tmpIndex] {
				maxArrayValue = tmpArray[tmpIndex]
			}
		}
		fmt.Printf("%v: %v\n", value, myArray)
	}

	return maxArrayValue

}
