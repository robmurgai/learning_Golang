package hackerrank

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Constraints
//		1 ≤ q ≤ 10^5, q is the number of queries
//		1 ≤ x, y, z ≤ 10^9, x,y,z are the value
//		All Querries [i][0] belong to {1, 2, 3}
//
// Example querries = [(1,1)(2,2)(3,2)(1,1)(1,1)(2,1)(3,2)]
// 		The results of each operation are:
// 		Operation   Array   dsOfFrequencies
// 		(1,1)       [1]
// 		(2,2)       [1]
// 		(3,2)                   0
// 		(1,1)       [1,1]
// 		(1,1)       [1,1,1]
// 		(2,1)       [1,1]
// 		(3,2)                   1
// Returns an array with the dsOfFrequencies: [0,1]
// Input 0
// 		{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}
// dsOfFrequencies 0
//		[0, 1]
// Explanation 0
// 		For the first query of type 3, there is no integer whose frequency is  2(array = [5,6]). So answer is 0.
//		For the second query of type 2, there are two integers in array=[6,19,10,6] whose frequency is 2 (integers 6 and 10). So, the answer is 1.
// Input 1
//		{3, 4}, {2, 1003}, {1, 16}, {3, 1}
// dsOfFrequencies 1
//		[0, 1]
// Explanation 1
// 		For the first query of type 3, there is no integer of frequency 4. The answer is 0. For the second query of type 3, there is one integer,16
// 		of frequency 1 so the answer is 1.
// Input 2
//		{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}
// dsOfFrequencies 2
//		[0, 1, 1]
// Explanation 2
// 		When the first dsOfFrequencies query is run, the array is empty. We insert two 4's and two 5's before the second dsOfFrequencies query, arr=[4,5,5,4] so there are two
//		instances of elements occurring twice. We delete a 4 and run the same query. Now only the instances of 5 satisfy the query.

// FreqQuery given a 2-d array of querries of 3 types, returns the results of queries of type 3
// Queries are given in the form of a 2-D array of size q where q[i][0] contains the operation, and q[i][1] contains the data element.
// Each query is of the form two integers described below:
// 		(1,x) : Insert x in your data structure.
// 		(2,y) : Delete one occurence of y from your data structure, if present.
// 		(3,z) : Check if any integer is present whose frequency is exactly z. If yes, print 1 else 0.
// Returns
// 		int[]: the results of queries of type 3
func DebugFreqQuery(queries [][]int32) []int32 {

	ans := createQueryOutput()

	type debugfqDS struct {
		dbl      []int32
		response int32
		ans      int32
	}

	debugfq := make([]debugfqDS, 0, len(queries))

	var response []int32

	// Create two variables,
	// A map where map.key is the number and map.value is the frequency of that number, hence adding and deleting a data occurance is constant.
	dsOfNumbers := make(map[int32]int32)

	// A map where map.key was the frequency of a number and map.value was the number, hence finding a number with a certain frequency was constant.
	dsOfFrequencies := make(map[int32]int32)

	for _, dbl := range queries {
		query := dbl[0]

		if query == int32(1) { // (1,x) : Insert x in your data structure.

			x := dbl[1]
			var curFreqOfx int32
			var newFreqOfx int32

			// Set current frequence of x
			//if this is first occurence of x, initialize data structure of numbers for x to 0.
			if _, isPresent := dsOfNumbers[x]; !isPresent {
				dsOfNumbers[x] = 0
			}
			curFreqOfx = dsOfNumbers[x]

			// Decreament the currnet frquency of x by 1, unless its currently 0.
			// if this is the first occurance of this frequency, initialize it to 0
			if _, isPresent := dsOfFrequencies[curFreqOfx]; !isPresent {
				dsOfFrequencies[curFreqOfx] = 0
			}
			if dsOfFrequencies[curFreqOfx] != 0 {
				dsOfFrequencies[curFreqOfx]--
			}

			//increment the frequency of x by 1
			dsOfNumbers[x]++

			// Incrementing the new frquency of x by 1
			newFreqOfx = dsOfNumbers[x]

			// if this is the first occurance of this particular frequency, initialize data structure of frequcenies to 0
			if _, isPresent := dsOfFrequencies[newFreqOfx]; !isPresent {
				dsOfFrequencies[newFreqOfx] = 0
			}
			dsOfFrequencies[newFreqOfx]++

			//fmt.Printf("(%v %-2v): Add %v to dsOfNumbers: %v.  dsOfFrequencies: %v\n", query, x, x, dsOfNumbers, dsOfFrequencies)

		} else if query == int32(2) {
			y := dbl[1]
			var curFreqOfy int32
			var newFreqOfy int32

			if _, isPresent := dsOfNumbers[y]; isPresent {
				curFreqOfy = dsOfNumbers[y]
				if dsOfFrequencies[curFreqOfy] != 0 {
					dsOfFrequencies[curFreqOfy]--
				}

				if dsOfNumbers[y] != 0 {
					dsOfNumbers[y]--
				}

				newFreqOfy = dsOfNumbers[y]
				// if this is the first occurance of this particular frequency, initialize data structure of frequcenies to 0
				if _, isPresent := dsOfFrequencies[newFreqOfy]; !isPresent {
					dsOfFrequencies[newFreqOfy] = 0
				}
				dsOfFrequencies[newFreqOfy]++
			}
			//fmt.Printf("(%v %-2v): Remove %v from dsOfNumbers: %v.  dsOfFrequencies: %v\n", query, y, y, dsOfNumbers, dsOfFrequencies)
		} else if query == int32(3) {
			z := dbl[1]
			printValue := 0
			if numAtFreqZ, isPresent := dsOfFrequencies[z]; isPresent {
				if numAtFreqZ > 0 {
					printValue = 1
				}
			}
			response = append(response, int32(printValue))

			//fmt.Printf("(%v %-2v) Check if any integers with frequencey %v are present in dsOfFrequencies: %v.  If so Update Response: %v\n", query, z, z, dsOfFrequencies, response)
		} else {
			// fmt.Printf("%v: Query Type: %v not recognized\n", dbl, query)
		}

		//Debug check if it matches the answer
		var temp debugfqDS
		var printMe = false
		temp.dbl = dbl
		if query == int32(3) {
			temp.response = response[len(response)-1]
			temp.ans = ans[len(response)-1]
			if temp.response != temp.ans {
				printMe = true
			}
		} else {
			if dsOfNumbers[temp.dbl[1]] < 0 {
				fmt.Printf("Debug: -ve Frequency Found, dsOfNumbers[%v]: %v\n", temp.dbl[1], dsOfNumbers[temp.dbl[1]])
				printMe = true
			}

		}
		debugfq = append(debugfq, temp)
		if printMe {
			fmt.Printf("Debug: Values don't match: %+v\n", temp)
			fmt.Printf("Debug: dsOfNumbers: %+v\n", dsOfNumbers)
			fmt.Printf("Debug: dsOfFrequencies: %+v\n", dsOfFrequencies)
			break
		}

	}
	//fmt.Printf("\nResponse%v\n", response)

	return response
}

func FreqQuery(queries [][]int32) []int32 {

	var response []int32

	// Create two variables,
	// A map where map.key is the number and map.value is the frequency of that number, hence adding and deleting a data occurance is constant.
	dsOfNumbers := make(map[int32]int32)

	// A map where map.key was the frequency of a number and map.value was the number, hence finding a number with a certain frequency was constant.
	dsOfFrequencies := make(map[int32]int32)

	for _, dbl := range queries {
		query := dbl[0]

		if query == int32(1) { // (1,x) : Insert x in your data structure.

			x := dbl[1]
			var curFreqOfx int32
			var newFreqOfx int32

			// Set current frequence of x
			//if this is first occurence of x, initialize data structure of numbers for x to 0.
			if _, isPresent := dsOfNumbers[x]; !isPresent {
				dsOfNumbers[x] = 0
			}
			curFreqOfx = dsOfNumbers[x]

			// Decreament the currnet frquency of x by 1, unless its currently 0.
			// if this is the first occurance of this frequency, initialize it to 0
			if _, isPresent := dsOfFrequencies[curFreqOfx]; !isPresent {
				dsOfFrequencies[curFreqOfx] = 0
			}
			if dsOfFrequencies[curFreqOfx] != 0 {
				dsOfFrequencies[curFreqOfx]--
			}

			//increment the frequency of x by 1
			dsOfNumbers[x]++

			// Incrementing the new frquency of x by 1
			newFreqOfx = dsOfNumbers[x]

			// if this is the first occurance of this particular frequency, initialize data structure of frequcenies to 0
			if _, isPresent := dsOfFrequencies[newFreqOfx]; !isPresent {
				dsOfFrequencies[newFreqOfx] = 0
			}
			dsOfFrequencies[newFreqOfx]++

		} else if query == int32(2) {
			y := dbl[1]
			var curFreqOfy int32
			var newFreqOfy int32

			if _, isPresent := dsOfNumbers[y]; isPresent {
				curFreqOfy = dsOfNumbers[y]
				if dsOfFrequencies[curFreqOfy] != 0 {
					dsOfFrequencies[curFreqOfy]--
				}

				if dsOfNumbers[y] != 0 {
					dsOfNumbers[y]--
				}

				newFreqOfy = dsOfNumbers[y]
				// if this is the first occurance of this particular frequency, initialize data structure of frequcenies to 0
				if _, isPresent := dsOfFrequencies[newFreqOfy]; !isPresent {
					dsOfFrequencies[newFreqOfy] = 0
				}
				dsOfFrequencies[newFreqOfy]++
			}
		} else if query == int32(3) {
			z := dbl[1]
			printValue := 0
			if numAtFreqZ, isPresent := dsOfFrequencies[z]; isPresent {
				if numAtFreqZ > 0 {
					printValue = 1
				}
			}
			response = append(response, int32(printValue))
		}
	}
	return response
}

func createQueryOutput() []int32 {

	// Read the input File
	dir := "data_files"
	fileName := "findQuery_output08.txt"
	filePath := filepath.Join(dir, fileName)
	//fmt.Printf("DEBUG: filePath: %v\n", filePath)

	//Intialize a buffer for the file
	buf, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: " + err.Error())
	}
	defer buf.Close()

	//Read the file line by line
	scannedLine := bufio.NewScanner(buf)

	var response []int32

	breakScanner := 1

	for scannedLine.Scan() {

		queriesRowString := scannedLine.Text()
		//fmt.Printf("Scanned Line(%v): %v\n", breakScanner, queriesRowString)

		// Do things with the scanned Line.
		queriesItem64, err := strconv.ParseInt(queriesRowString, 10, 64)
		checkError(err)
		queriesItem := int32(queriesItem64)

		response = append(response, queriesItem)

		breakScanner++
		// if breakScanner == 10 {
		// 	break
		// }
	}

	err = scannedLine.Err()
	if err != nil {
		fmt.Printf("Error: " + err.Error())
	}

	return response
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
