/**
 **
 ** Experiment: words.go
 ** Write a function to count the frequency of words in a string of text and return a map of words with their counts. The function should convert
 ** the text to lowercase, and punctuation should be trimmed from words. The strings package contains several helpful functions for this task,
 ** including Fields, ToLower, and Trim.
 **
 ** Use your function to count the frequency of words in the following passage and then display the count for any word that occurs more than once.
 **
 ** As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple
 ** transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again;
 ** the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once
 ** or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the
 ** fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.
 **
 **
 ** C.S. Lewis, Out of the Silent Planet, (see mng.bz/V7nO)
 **
 **/

package main

import (
	"fmt"
	"strings"
)

func words(sentence string) map[string]int {

	// Declare a map of words with their counts that we will return from this function
	mapOfWordsWithTheirCounts := make(map[string]int)

	// The function should convert the text to lowercase, and punctuation should be trimmed from words. The strings package contains several
	// helpful functions for this task, including Fields, ToLower, and Trim.

	//Print the original sentence argument
	fmt.Printf("Original sentence passed as the argument: %v\n\n", sentence)

	//convert the argument to lower case
	sentence = strings.ToLower(sentence)
	//fmt.Printf("sentence after ToLower: %v\n\n", sentence)

	//Create a full slice from the string
	sentenceSlice := strings.Fields(sentence)
	//fmt.Printf("sentenceSlice: %v\n\n", sentenceSlice)

	for i, wordInSlice := range sentenceSlice {

		//trim punctuations in the slice
		sentenceSlice[i] = strings.Trim(wordInSlice, ". , ;")
		//fmt.Printf("After Trim sentenceSlice[%v]: %v\n", i, sentenceSlice[i])

		//Update the count for each word in the slice
		mapOfWordsWithTheirCounts[sentenceSlice[i]] = mapOfWordsWithTheirCounts[sentenceSlice[i]] + 1
		//fmt.Printf("This is occurance %v of the word \"%v\" in the sentence\n\n", mapOfWordsWithTheirCounts[sentenceSlice[i]], sentenceSlice[i])

	}

	for key, value := range mapOfWordsWithTheirCounts {
		if value > 1 {
			fmt.Printf("The word \"%v\" occurs %v times\n", key, value)
		}

	}
	//fmt.Printf("The map of words with their counts: %v", mapOfWordsWithTheirCounts)

	return mapOfWordsWithTheirCounts
}
