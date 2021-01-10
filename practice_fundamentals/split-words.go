// Practice starting a goroutine, using channels to communicate and understanding channel pipelines
// split-words.go
// Sometimes it’s easier to operate on words than on sentences. Write a pipeline element that takes strings, splits them up into words (you can use
// the Fields function from the strings package), and sends all the words, one by one, to the next pipeline stage.
//
// Use the following main function
// func main() {

// 	fmt.Printf("########## Starting Exercise ##########\n\n")

// 	downstream := make(chan string)

// 	go remIden(downstream)

// 	for _, v := range []string{"This is a string of 7 words"} {
// 		fmt.Println("Calling splitWords('", v, "')")
// 		splitWords(v, downstream)
// 	}

// 	time.Sleep(10)
// 	close(downstream)
// 	fmt.Printf("\n########## Ending Exercise ##########\n")

// }

package practicefundamentals

import (
	"fmt"
	"strings"
)

func splitWords(s string, downstream chan string) {

	for _, v := range strings.Fields(s) {
		fmt.Println("Sending: ", v)
		downstream <- v
	}
}
