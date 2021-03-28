package main

import (
	"fmt"
)

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	//Unsorted List
	var s string
	var numOfAnnagrams int32

	// s = "bba"
	// fmt.Printf("String s: %v sorted: %v", s, mergeSort(s))

	s = "abba"
	fmt.Printf("%v: %v\n", s, getSignature(s))
	fmt.Printf("string: %v, should have 4 annagrams\n", s)
	numOfAnnagrams = sherlockAndAnagrams(s)
	fmt.Printf("%v has %v annagrams\n", s, numOfAnnagrams)

	s = "abcd"
	fmt.Printf("string: %v should have 0 annagrams\n", s)
	numOfAnnagrams = sherlockAndAnagrams(s)
	fmt.Printf("%v has %v annagrams\n", s, numOfAnnagrams)

	s = "kkkk"
	fmt.Printf("string: %v should have 10 annagrams\n", s)
	numOfAnnagrams = sherlockAndAnagrams(s)
	fmt.Printf("%v has %v annagrams\n", s, numOfAnnagrams)

	//Partially Sorted List
	//Sorted List

	fmt.Printf("\n########## Ending Exercise ##########\n")
}

func sherlockAndAnagrams(s string) int32 {
	var bigO = 0
	var response int32 = 0

	// string: abba, abcd
	// Case 1: no matching Chars in the string, return 0
	// Case 2: matching Chars
	//    Create substrings
	//    Sort substrings
	//    Store all sorted substrings in a map.
	//    return the number of matching substrings

	// Case 1: no matching Chars in the string, return 0
	sherlock := make(map[string]int)
	for _, sChars := range s {
		bigO++
		sub := string(sChars)
		subString := fmt.Sprint(getSignature(sub))
		sherlock[subString]++
		if sherlock[subString] > 1 {
			// fmt.Printf(" SandA: Found %v string %v times\n", subString, sherlock[subString])
			response += int32(sherlock[subString] - 1)
			// fmt.Printf(" SandA: Updating # of pairs to: %v\n", response)
		}
	}
	if response == 0 {
		// fmt.Printf(" SandA: For String: %v\n", s)
		// fmt.Printf(" SandA: sherlock: %v\n", sherlock)
		// fmt.Printf(" SandA: %v has no matching chars\n", s)
		// fmt.Printf(" SandA: String size: %v, bigO: %v\n", len(s), bigO)
		return response
	}

	// Case 2: matching Chars
	//    Create substrings
	//    Sort substrings
	//    Store all sorted substrings in a map.
	//    return the number of matching substrings

	//substring size goes from n-1 to 2
	for substringSize := len(s) - 1; substringSize > 1; substringSize-- {
		for leftIndex := 0; leftIndex+substringSize <= len(s); leftIndex++ {
			bigO++
			sub := s[leftIndex : leftIndex+substringSize]
			// fmt.Printf(" SandA: subString: %v\n", subString)
			subString := fmt.Sprint(getSignature(sub))
			// fmt.Printf(" SandA: sorted subString: %v\n", subString)
			sherlock[subString]++
			if sherlock[subString] > 1 {
				// fmt.Printf(" SandA: Found %v string %v times\n", subString, sherlock[subString])
				response += int32(sherlock[subString] - 1)
				// fmt.Printf(" SandA: Updating # of pairs to: %v\n", response)
			}
		}
	}

	// fmt.Printf(" SandA: For String: %v\n", s)
	// fmt.Printf(" SandA: sherlock: %v\n", sherlock)
	// fmt.Printf(" SandA: String size: %v, bigO: %v\n", len(s), bigO)
	return response
}

func getSignature(s string) []int {
	signature := make([]int, 26)

	for _, r := range s {
		signature[int(r)-97]++
	}

	return signature
}

func sherlockAndAnagrams1(s string) int32 {
	var bigO = 0
	var response int32 = 0

	// string: abba, abcd
	// Case 1: no matching Chars in the string, return 0
	// Case 2: matching Chars
	//	Create substrings
	//	Sort substrings
	//	Store all sorted substrings in a map.
	//	return the number of matching substrings

	// Case 1: no matching Chars in the string, return 0

	sherlock := make(map[string][]int)
	magicIndexes := make(map[int]bool)

	for index, sChars := range s {
		bigO++
		subString := string(sChars)
		if _, isPresent := sherlock[subString]; isPresent {
			sherlock[subString][0]++
		} else {
			sherlock[subString] = append(sherlock[subString], 1)
		}
		sherlock[subString] = append(sherlock[subString], index)
		magicIndexes[index] = true

		if sherlock[subString][0] > 1 {
			fmt.Printf(" SandA: Found %v string %v times\n", subString, sherlock[subString])
			response += int32(sherlock[subString][0] - 1)
			fmt.Printf(" SandA: Updating # of pairs to: %v\n", response)
		}
	}
	if response == 0 {
		fmt.Printf(" SandA: For String: %v\n", s)
		fmt.Printf(" SandA: sherlock: %v\n", sherlock)
		fmt.Printf(" SandA: %v has no matching chars\n", s)
		fmt.Printf(" SandA: String size: %v, bigO: %v\n", len(s), bigO)
		return response
	}

	// Case 2: matching Chars
	//	Create substrings
	//	Sort substrings, if they have matching chars.
	//	Store all sorted substrings in a map.
	//	return the number of matching substrings

	//substring size goes from n-1 to 2
	for substringSize := len(s) - 1; substringSize > 1; substringSize-- {
		for leftIndex := 0; leftIndex+substringSize <= len(s); leftIndex++ {
			bigO++
			subString := s[leftIndex : leftIndex+substringSize]
			fmt.Printf(" SandA: subString: %v\n", subString)

			// if substring has a matching char, sort it, and do the rest, othewise, move on
			subString = mergeSort(subString)

			fmt.Printf(" SandA: sorted subString: %v\n", subString)

			if _, isPresent := sherlock[subString]; isPresent {
				sherlock[subString][0]++
			} else {
				sherlock[subString] = append(sherlock[subString], 1)
			}
			sherlock[subString] = append(sherlock[subString], leftIndex)
			if sherlock[subString][0] > 1 {
				fmt.Printf(" SandA: Found %v string %v times\n", subString, sherlock[subString])
				response += int32(sherlock[subString][0] - 1)
				fmt.Printf(" SandA: Updating # of pairs to: %v\n", response)
			}

		}
	}

	fmt.Printf(" SandA: For String: %v\n", s)
	fmt.Printf(" SandA: sherlock: %v\n", sherlock)
	fmt.Printf(" SandA: String size: %v, bigO: %v\n", len(s), bigO)
	return response
}

func mergeSort(s string) string {

	fmt.Printf("\n  mergeSort: %v\n", s)

	// Case 1: If s is a single char, return it, its already sorted.
	// Case 2: If s is 2 or more chars, break it in 2 strings and call merge sort on both
	//   Take the retruned strings and merge them.

	var response string = s
	if len(s) <= 1 {
		fmt.Printf("  mergeSort: sorted string: %v\n\n", s)
		return s
	}

	leftString := s[:len(s)/2]
	fmt.Printf("  mergeSort: %v: leftString: %v\n", s, leftString)
	leftString = mergeSort(leftString)
	fmt.Printf("  mergeSort: %v: sorted leftString: %v\n", s, leftString)

	rightString := s[len(s)/2:]
	fmt.Printf("  mergeSort: %v: rightString: %v\n", s, rightString)
	rightString = mergeSort(rightString)
	fmt.Printf("  mergeSort: %v: sorted rightString: %v\n", s, rightString)

	response = mergeStrings(leftString, rightString)

	fmt.Printf("  mergeSort: %v: sorted string: %v\n\n", s, response)
	return response
}

func mergeStrings(leftString, rightString string) string {

	fmt.Printf("  mergeStrings: %v & %v\n", leftString, rightString)
	var response string = ""

	//Merge left string charectors into right string charectors
	//Case 1: leftStringChar is <= rightStringChar, append leftStringChar to response
	//Case 2: leftStringChar is > rightStringChar, append rightStringChar to response
	//Case 3: If all the right string chars have been appended, append the rest of the leftString
	//Cae 4: If all the left string charts have been appended, append the rest of the right string

	//b and ab
	rightStringIndex := 0

	for leftStringIndex := 0; leftStringIndex < len(leftString); leftStringIndex++ {
		fmt.Printf("  mergeStrings: leftStringIndex: %v\n", leftStringIndex)
		fmt.Printf("  mergeStrings: rightStringIndex: %v\n", rightStringIndex)

		leftStringChar := leftString[leftStringIndex]

		for rightStringIndex < len(rightString) {
			rightStringChar := rightString[rightStringIndex]
			fmt.Printf("  mergeStrings: Comparing %c & %c\n", leftStringChar, rightStringChar)
			if leftStringChar <= rightStringChar {
				fmt.Printf("  mergeStrings: %c <= %c\n", leftStringChar, rightStringChar)
				response = response + string(leftStringChar)
				fmt.Printf("  mergeStrings: response: %v\n", response)
				break
			}

			fmt.Printf("  mergeStrings: %c > %c\n", leftStringChar, rightStringChar)
			response = response + string(rightStringChar)
			fmt.Printf("  MmergeStrings: response: %v\n", response)
			rightStringIndex++
		}

		if rightStringIndex == len(rightString) && leftStringIndex < len(leftString) {
			response = response + string(leftString[leftStringIndex:])
			fmt.Printf("  MmergeStrings: response: %v\n", response)
		}

	}

	if rightStringIndex < len(rightString) {
		response = response + string(rightString[rightStringIndex:])
	}

	fmt.Printf("  MmergeStrings: response: %v\n", response)
	return response

}
