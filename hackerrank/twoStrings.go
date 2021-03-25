// Interview Preparation Kit | Dictionaries and Hashmaps | Two Strings
// Given two strings, determine if they share a common substring. A substring may be as small as one character.

// Example

// These share the common substring .

// These do not share a substring.

// Function Description

// Complete the function twoStrings in the editor below.

// twoStrings has the following parameter(s):

// string s1: a string
// string s2: another string
// Returns

// string: either YES or NO
// Input Format

// The first line contains a single integer , the number of test cases.

// The following  pairs of lines are as follows:

// The first line contains string .
// The second line contains string .
// Constraints

//  and  consist of characters in the range ascii[a-z].
// Output Format

// For each pair of strings, return YES or NO.

// Sample Input

// 2
// hello
// world
// hi
// world
// Sample Output

// YES
// NO
// Explanation

// We have  pairs to check:

// , . The substrings  and  are common to both strings.
// , .  and  share no common substrings.

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {

	response := "NO"

	//hello world
	//hi world

	//Steps:
	// hello
	// single char string: h e l l o, find them in string 2nd string world, 'l' will be found

	// Optimized Search Strategies
	// searches are better with maps: so converting a string to a map would be faster?
	//      if the 2nd string was a map[string]bool{"w": true, "o": true, "r": true, ..}
	//      find "hello" in the map, map["h"]: exist?
	// Do we need to create a map of all lenght strings, single, 2 chars, 3 chars, etc
	//      No: if no single char string found, then no mulitple char string would be found?
	// Should we start with single chart strings and move up in size, or full string and move down
	//      Single Char start is O(n) iterations to build the map + O(1) to find
	//      Multiple Char start is O(n!) + O(1) to find the string.

	// Convert string2 into a map
	var s2Map map[rune]bool = make(map[rune]bool, len(s2))
	for _, key := range s2 {
		s2Map[key] = true
	}

	//check if any single char from string 1 is present in string 2
	for _, value := range s1 {
		if _, isPresent := s2Map[value]; isPresent {
			fmt.Printf("Char %c from string %v is present in string %v\n", value, s1, s2)
			response = "YES"
			break
		}
	}
	return response

}