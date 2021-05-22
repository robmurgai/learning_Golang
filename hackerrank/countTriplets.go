package hackerrank

import "fmt"

// You are given an array and you need to find number of tripets of indices such that the elements
// at those indices are in geometric progression for a given common ratio and .
// For example, arr = [1, 4, 16, 64]. If r=4, we have [1,4,16] and [4,16,64] at indices (0,1,2) and (1,2,3)
//
// Function Description
// Complete the countTriplets function in the editor below. It should return the number of triplets forming a
// geometric progression for a given r as an integer.
// countTriplets has the following parameter(s):
// 		arr: an array of integers
// 		r: an integer, the common ratio
//
// Input Format
// The first line contains two space-separated integers n and r, the size of arr and the common ratio.
// The next line contains n space-seperated integers arr[i].
// Constraints
//		1 <= n <= 10^5
//		1 <= r <= 10^9
//		1 <= arr[i] <= 10^9
//
// Output Format
// Return the count of triplets that form a geometric progression.

// Complete the countTriplets function below.
func CountTriplets1(arr []int64, r int64) int64 {

	// For every number in the array, its either divisible by r, or its not.
	// If it is divisible by r, take the divident and multiplier and see if it exists in the array.
	// Triplets are (i, j, k)
	// {1,2,2,4} with r=2, has 2 triplets (0,1,3) and (0,2,3)
	// {1:{0}, 2:{1,2}, 4:{3}}
	// 1: No, ignore
	// 2: Yes at index 1, Divident 1 at Index 0, Multiplier 4 at index 3 --> i=0, j=1, k=3, Num of triplets 1*1*1 = 1
	// 2: Yes at index 2, Divident 1, Divident Index 0, Multiplier 4 at index 3 --> i=0, j=1, k=3, Num of triplets 1*1*1 = 1
	// 4: Yes at index 3, Divident 2 at Index 1, 2; Multiplier 8 not found -->  dropped.
	//
	// {1,3,9,9,27,81} with r=3 has 6 triplets (0,1,2) (0,1,3) (1,2,4) (1,3,4)(2,4,5) and
	// (3,4,5)
	// {1:{0}, 3:{1}, 9:{2,3}, 27{4}, 81{5}}
	// 1: No, ignore
	// 3: Yes at 1, divident 1 at index 0, Multiplier 9 at index 2,3 --> Num of triplets 1*1*2 = 2
	// 9: Yes at 2, divident 3 at index 1, Multiplier 27 at index 4 --> Numb of triplets 1*1*1 = 1
	// 9: Yes at 3, divident 3 at index 1, Multiplier 27 at index 4 --> Numb of triplets 1*1*1 = 1
	// 27: Yes at 4, divident 9 at index 2, 3; Multiplier 81 at index 5 --> Numb of triplets 1*2*1 = 2
	// 81: Yes at 5, divident 27 at index 4, Mulitplier 273 not found ---> dropped.

	fmt.Println(arr)
	fmt.Println(r)
	var numOfTrips int64
	fmt.Printf("numOfTiplets: %v\n", numOfTrips)

	foo := make(map[int64][]int, len(arr))

	for index, value := range arr {
		foo[value] = append(foo[value], index)
	}

	fmt.Printf("foo: %v\n", foo)

	for index, value := range arr {

		if value%r == 0 {
			//j
			midTrip := 1
			fmt.Printf("j found: [%v]\n", index)

			//i
			divident := value / r
			if dValue, isPresent := foo[divident]; isPresent {
				fmt.Printf("i found: %v\n", dValue)
				leftTrip := len(dValue)
				//k
				multiplier := value * r
				if mValue, isPresent := foo[multiplier]; isPresent {
					fmt.Printf("k found: %v\n", mValue)
					fmt.Printf("triplet: (%v, [%v], %v)\n", dValue, index, mValue)

					rightTrip := len(mValue)
					numOfTrips += int64(midTrip * leftTrip * rightTrip)
					fmt.Printf("numOfTiplets incremented to: %v\n", numOfTrips)

				}
			}
		}
	}
	return numOfTrips
}

// Complete the countTriplets function below.
func CountTriplets2(arr []int64, r int64) int64 {

	// For every number in the array, its either divisible by r, or its not.
	// If it is divisible by r, take the divident and multiplier and see if it exists in the array.
	// Triplets are (i, j, k)
	// {1,2,2,4} with r=2, has 2 triplets (0,1,3) and (0,2,3)
	// {1:{0}, 2:{1,2}, 4:{3}}
	// 1: No, ignore
	// 2: Yes at index 1, Divident 1 at Index 0, Multiplier 4 at index 3 --> i=0, j=1, k=3, Num of triplets 1*1*1 = 1
	// 2: Yes at index 2, Divident 1, Divident Index 0, Multiplier 4 at index 3 --> i=0, j=1, k=3, Num of triplets 1*1*1 = 1
	// 4: Yes at index 3, Divident 2 at Index 1, 2; Multiplier 8 not found -->  dropped.
	//
	// {1,3,9,9,27,81} with r=3 has 6 triplets (0,1,2) (0,1,3) (1,2,4) (1,3,4)(2,4,5) and
	// (3,4,5)
	// {1:{0}, 3:{1}, 9:{2,3}, 27{4}, 81{5}}
	// 1: No, ignore
	// 3: Yes at 1, divident 1 at index 0, Multiplier 9 at index 2,3 --> Num of triplets 1*1*2 = 2
	// 9: Yes at 2, divident 3 at index 1, Multiplier 27 at index 4 --> Numb of triplets 1*1*1 = 1
	// 9: Yes at 3, divident 3 at index 1, Multiplier 27 at index 4 --> Numb of triplets 1*1*1 = 1
	// 27: Yes at 4, divident 9 at index 2, 3; Multiplier 81 at index 5 --> Numb of triplets 1*2*1 = 2
	// 81: Yes at 5, divident 27 at index 4, Mulitplier 273 not found ---> dropped.

	fmt.Println(arr)
	fmt.Println(r)
	var numOfTrips int64
	fmt.Printf("numOfTiplets: %v\n", numOfTrips)

	foo := make(map[int64][]int, len(arr))

	for index, value := range arr {
		foo[value] = append(foo[value], index)
		fmt.Printf("foo: %v\n", foo)

		if value%r == 0 {
			//k
			//numOfkIndeces := 1
			fmt.Printf("k found: [%v]\n", index)

			//j
			midValue := value / r
			if midValueIndeces, isPresent := foo[midValue]; isPresent {
				fmt.Printf("j found: %v\n", midValueIndeces)
				//numOfjIndeces := len(midValueIndeces)

				//i
				lefValue := midValue / r
				if leftValueIndeces, isPresent := foo[lefValue]; isPresent {
					fmt.Printf("i found: %v\n", leftValueIndeces)
					//numOfiIndeces := len(leftValueIndeces)
					fmt.Printf("triplet: (%v, [%v], %v)\n", leftValueIndeces, midValueIndeces, index)

					//numOfTrips
					// for _, jIndex := range midValueIndeces {
					// 	for _, iIndex := range leftValueIndeces {
					// 		if iIndex < jIndex && jIndex < index {
					// 			numOfTrips++
					// 		}
					// 	}
					// }

					// smallest k > biggest J && smalles J > bigges i
					for _, iIndex := range leftValueIndeces {
						for j, jIndex := range midValueIndeces {
							if iIndex < jIndex {
								numOfTrips += int64(len(midValueIndeces) - j)
								break
							}
						}
					}
					fmt.Printf("numOfTiplets incremented to: %v\n", numOfTrips)

				}
			}
		}
	}

	return numOfTrips
}

// Complete the countTriplets function below.
func CountTriplets3(arr []int64, r int64) int64 {

	var numOfTrips int64

	type bar struct {
		positions    []int
		dbl          bool
		dblValue     int64
		dblPositions []int
	}

	foo := make(map[int64]*bar, len(arr))

	//Keep track of the number of divisers (aka doubles) to the left.

	// for each value in arr.
	// is it a double?
	// if so,
	//		mark it as such and
	//		store the number of dbls at that position
	//		and check is it a triple
	// if so, increment numOfTrips
	// Add current index.

	for index, value := range arr {

		//Lets stipulate, this is the k Value in Triplet (i < j < k)
		kValue := value
		//fmt.Printf("%v @ index %v could be a k value\n", kValue, index)

		// Add current index.
		var tmp bar
		if _, isPresent := foo[kValue]; isPresent {
			tmp = *foo[kValue]
		} else {
			tmp.positions = make([]int, 0)
		}

		tmp.positions = append(tmp.positions, index)

		//fmt.Println(foo[kValue])
		foo[kValue] = &tmp

		//Is k Value a double, does it have a devisor in the array.
		// if so,
		//		mark it as such and
		//		store the number of dbls at that position
		if kValue%r == 0 {
			jValue := kValue / r
			if _, isPresent := foo[jValue]; isPresent {
				foo[kValue].dbl = true
				foo[kValue].dblValue = jValue
				foo[kValue].dblPositions = foo[jValue].positions
				// dblPositions should be numOfDbls
				// we should append the len(foo[jValue].positions)
				// fmt.Printf("  %v @ index %v could be a j value\n", jValue, foo[kValue].dblPositions)

				//Is it a triple --> Is j Value a double, does it also have a devisor in the array.
				if foo[jValue].dbl {
					iValue := foo[jValue].dblValue
					fmt.Printf("%v @ index %v is an i value\n", iValue, foo[jValue].dblPositions)
					fmt.Printf("%v @ index %v is a j value\n", jValue, foo[kValue].dblPositions)
					fmt.Printf("%v @ index %v is a k value\n", kValue, index)

					// numOfTrips += int64(midTrip * leftTrip * rightTrip)

				}

			}
		}
		// fmt.Printf("Index: %v Value: %v\nfoo[%v] { \n  %+v\n}\n\n", index, value, kValue, foo[kValue])

		for index, value := range foo {
			fmt.Printf("foo[%v]: %+v \n", index, value)
		}
		fmt.Println()
	}

	return numOfTrips
}

// CountTriplets counts the number of triplets.
// Keep track of the number of divisers (aka doubles) to the left.
//
// for each value in arr.
// is it a double?
// if so,
//		mark it as such and
//		store the number of dbls at that position
//		and check is it a triple
// if so, increment numOfTrips
// Add current index.
func CountTriplets(arr []int64, r int64) int64 {

	var numOfTrips int64

	if r == 1 {

		// bar has positions, numOfTriples []int
		type bar struct {
			positions    []int
			numOftriples []int64
		}
		foo := make(map[int64]*bar, len(arr))

		for index, value := range arr {

			//fmt.Printf("For r: %v, %v @ index %v\n", r, value, index)

			// Add current index.
			var tmp bar
			if _, isPresent := foo[value]; isPresent {
				tmp = *foo[value]
			} else {
				tmp.positions = make([]int, 0)
				tmp.numOftriples = make([]int64, 0)
			}

			tmp.positions = append(tmp.positions, index)

			if len(tmp.positions) >= 3 {

				n := len(tmp.positions) - 2
				tripCount := int64((n * (n + 1)) / 2)

				tmp.numOftriples = append(tmp.numOftriples, tripCount)
				numOfTrips += int64(tripCount)
				// fmt.Printf("Index %v; Len %v; n: %v; Triplets %v. Total Triplets: %v\n", index, len(tmp.positions), n, tripCount, numOfTrips)
			}
			foo[value] = &tmp
		}
	} else {

		type bar struct {
			positions []int
			dbl       bool
			dblValue  int64
			numOfDbls []int
		}

		foo := make(map[int64]*bar, len(arr))

		for index, value := range arr {

			//fmt.Printf("%v @ index %v\n", kValue, index)

			//Lets stipulate, this is the k Value in Triplet (i < j < k)
			kValue := value

			// Add current index.
			var tmp bar
			if _, isPresent := foo[kValue]; isPresent {
				tmp = *foo[kValue]
			} else {
				tmp.positions = make([]int, 0)
			}

			tmp.positions = append(tmp.positions, index)
			foo[kValue] = &tmp

			//Is k Value a double, does it have a devisor in the array.
			// if so,
			//		mark it as such and
			//		store the number of dbls at that position
			if kValue%r == 0 {
				jValue := kValue / r
				if _, isPresent := foo[jValue]; isPresent {

					foo[kValue].dbl = true
					foo[kValue].dblValue = jValue

					// Count the number of possible Dbls
					numOfDbls := len(foo[jValue].positions)
					foo[kValue].numOfDbls = append(foo[kValue].numOfDbls, numOfDbls)

					//Is it a triple --> Is j Value a double, does it also have a devisor in the array.
					if foo[jValue].dbl {
						//iValue := foo[jValue].dblValue
						//oldTrips := numOfTrips

						for _, count := range foo[jValue].numOfDbls {
							numOfTrips += int64(count)
						}
						//fmt.Printf("%v triplets found for i < j < k sequence for values %v %v %v.  Total Triplets: %v\n", numOfTrips-oldTrips, iValue, jValue, kValue, numOfTrips)
					}

				}
			}
			// for index, value := range foo {
			// 	fmt.Printf("foo[%v]: %+v \n", index, value)
			// }
			// fmt.Println()
		}
	}

	return numOfTrips
}
