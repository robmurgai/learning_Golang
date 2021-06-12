package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	hr "github.com/robmurgai/learning_Golang/hackerrank"
)

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	// Input 0
	// 		{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}
	// Output 0
	//		[0, 1]
	queries := [][]int32{{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}}
	fmt.Printf("\n\nCalling FrequencyQuery with parm input 0: %v\n", queries)
	fmt.Printf("Returned Output: %v\n", hr.FreqQuery(queries))
	fmt.Printf("Expected Output: [0 1]\n")

	// Input 1
	//		{3, 4}, {2, 1003}, {1, 16}, {3, 1}
	// Output 1
	//		[0, 1]
	// Explanation 1
	// 		For the first query of type 3, there is no integer of frequency 4. The answer is 0. For the second query of type 3, there is one integer,16
	// 		of frequency 1 so the answer is 1.
	queries = [][]int32{{3, 4}, {2, 1003}, {1, 16}, {3, 1}}
	fmt.Printf("\n\nCalling FrequencyQuery with parminput 1: %v\n", queries)
	fmt.Printf("Returned Output: %v\n", hr.FreqQuery(queries))
	fmt.Printf("Expected Output: [0 1]\n")

	// Input 2
	//		{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}
	// Output 2
	//		[0, 1, 1]
	queries = [][]int32{{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}}
	fmt.Printf("\n\nCalling FrequencyQuery with parm input 2: %v\n", queries)
	fmt.Printf("Returned Output: %v\n", hr.FreqQuery(queries))
	fmt.Printf("Expected Output: [0 1 1]\n")

	// Constraints Tests
	// 1 ≤ q ≤ 10^5, q is the number of queries
	queries = [][]int32{{1, 3}} // response should be []
	fmt.Printf("\n\nCalling FrequencyQuery with parm input 3: %v\n", queries)
	fmt.Printf("Returned Output: %v\n", hr.FreqQuery(queries))
	fmt.Printf("Expected Output: []\n")

	// queries = make([][]int32, 1e5)
	// for index := int32(0); index < 1e5; index++ {
	// 	queries[index] = []int32{1, 1}
	// }
	// queries[1e5-1] = []int32{3, 99999} // response should be [1]

	// Constraints Tests
	// 1 ≤ x, y, z ≤ 10^9, x,y,z are the value
	queries = make([][]int32, 10)
	for index := int32(0); index < 9; index++ {
		queries[index] = []int32{1, index + 999999992}
	}
	queries[9] = []int32{3, 1} // response should be [1]
	fmt.Printf("\n\nCalling FrequencyQuery with parm input 4: %v\n", queries)
	fmt.Printf("Returned Output: %v\n", hr.FreqQuery(queries))
	fmt.Printf("Expected Output: [1]\n")

	queries = [][]int32{{1, 1}, {1, 2}, {1, 1}, {3, 2}, {2, 1}, {3, 2}} // response should be [1, 0]
	fmt.Printf("\n\nCalling FrequencyQuery with parm input 5: %v\n", queries)
	fmt.Printf("Returned Output: %v\n", hr.FreqQuery(queries))
	fmt.Printf("Expected Output: [1 0]\n")

	queries = createQueryInput()
	solution := createQueryOutput()
	fmt.Printf("\n\nCalling FrequencyQuery with parm input 6 of length: %v\n", len(queries))
	fmt.Printf("Received Expected Output: %v", checkQueryAnswer(hr.FreqQuery(queries), solution))

	fmt.Printf("\n########## Ending Exercise ##########\n")
}

func createQueryInput() [][]int32 {

	// Read the input File
	dir := "data_files"
	fileName := "findQuery_input08.txt"
	filePath := filepath.Join(dir, fileName)
	// fmt.Printf("DEBUG: filePath: %v\n", filePath)

	//Intialize a buffer for the file
	buf, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: " + err.Error())
	}
	defer buf.Close()

	//Read the file line by line
	scannedLine := bufio.NewScanner(buf)

	//The first line is the size of the input array and should be handled seperately.
	scannedLine.Scan()
	querySizeString := scannedLine.Text()
	// fmt.Println("First Scanned Line: " + querySizeString)
	querySizeInt, err := strconv.ParseInt(querySizeString, 10, 64)
	checkError(err)
	queries := make([][]int32, 0, querySizeInt)

	breakScanner := 1

	for scannedLine.Scan() {

		queriesRowString := scannedLine.Text()
		//fmt.Printf("Scanned Line(%v): %v\n", breakScanner, queriesRowString)

		queriesRowSlice := strings.Split(queriesRowString, " ")
		//fmt.Printf("Scanned Line(%v): %v\n", breakScanner, queriesRowSlice)

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowSlice {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			fmt.Printf("Error: Unable to translate the line into a 2 index query: %v", queriesRow)
			panic("Error: Unable to translate the line into a 2 index query")
		}

		queries = append(queries, queriesRow)

		breakScanner++
		// if breakScanner == 10 {
		// 	break
		// }
	}
	err = scannedLine.Err()
	if err != nil {
		fmt.Printf("Error: " + err.Error())
	}
	return queries
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

func checkQueryAnswer(response []int32, ans []int32) bool {

	works := true

	//match the size
	if len(response) == len(ans) {

		//match individual values
		for index, responseValue := range response {
			if responseValue != ans[index] {
				works = false
				fmt.Printf("Values dont Match.  Reponse[%v]: %v, but expected Answer[%v]: %v\n", index, responseValue, index, ans[index])
				break
			}

		}
	} else {
		works = false
		fmt.Printf("The two lengths dont Match.  Reponse Size is: %v, but expected size was: %v\n", len(response), len(ans))
	}

	return works

}
