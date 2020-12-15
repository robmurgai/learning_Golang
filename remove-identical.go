// Practice starting a goroutine, using channels to communicate and understanding channel pipelines
// It’s boring to see the same line repeated over and over again. Write a pipeline element (a goroutine) that remembers the previous value and only
// sends the value to the next stage of the pipeline if it’s different from the one that came before. To make things a little simpler, you may assume
// that the first value is never the empty string.
//
// Use the following main function to call this solution
// func main() {

// 	fmt.Printf("########## Starting Exercise ##########\n\n")

// 	downstream := make(chan string)

// 	go remIden(downstream)

// 	for _, v := range []string{"This is 1", "This is 2", "This is 2", "This is 3"} {
// 		fmt.Println("Sending: ", v)
// 		downstream <- v
// 	}

// 	time.Sleep(10)
// 	close(downstream)
// 	fmt.Printf("\n########## Ending Exercise ##########\n")

// }
package main

import "fmt"

var prevString string

func remIden(downstream chan string) {
	fmt.Println("Starting remIden() goroutine")

	for curString := range downstream {
		fmt.Println("Received: ", curString)
		if curString != prevString {
			prevString = curString
			fmt.Println("Ready to send: ", curString)
		} else {
			fmt.Println("Both current and previous strings were the same: ", curString)
		}
	}
}
