// Two strings are anagrams of each other if the letters of one string can be rearranged to form the other string. Given a string, find
// the number of pairs of substrings of the string that are anagrams of each other.
//
// Sample main.go
// var s string
// var numOfAnnagrams int32

// s = "abba"
// fmt.Printf("string: %v, should have 4 annagrams\n", s)
// numOfAnnagrams = sherlockAndAnagrams(s)
// fmt.Printf("%v has %v annagrams\n", s, numOfAnnagrams)

// s = "abcd"
// fmt.Printf("string: %v should have 0 annagrams\n", s)
// numOfAnnagrams = sherlockAndAnagrams(s)
// fmt.Printf("%v has %v annagrams\n", s, numOfAnnagrams)

// s = "kkkk"
// fmt.Printf("string: %v should have 10 annagrams\n", s)
// numOfAnnagrams = sherlockAndAnagrams(s)
// fmt.Printf("%v has %v annagrams\n", s, numOfAnnagrams)

package hackerrank

import "fmt"

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
