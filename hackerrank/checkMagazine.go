package hackerrank

import "fmt"

// CheckMagazine must print Yes if the note can be formed using the magazine, or No.
// checkMagazine has the following parameters:
//
// magazine: an array of strings, each a word in the magazine
// note: an array of strings, each a word in the ransom note
//
// Constriants
// len(magaxine) and len(note) is less than 30000
// len(magaxine[i]) and len(note[i]) is less than 6
//
func CheckMagazine(magazine []string, note []string) {

	//convert magazine to a map
	magMap := make(map[string]int, len(magazine))
	for _, value := range magazine {
		magMap[value]++
	}

	for _, value := range note {
		wordCount, ok := magMap[value]
		if ok {
			if wordCount > 1 {
				magMap[value]--
			} else {
				//delete the key to avoid double count
				delete(magMap, value)
			}
		} else {
			//fmt.Printf("DEBUG: Note Word: %s not found in Magazine\n", value)
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
