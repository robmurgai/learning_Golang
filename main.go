package main

import (
	"fmt"

	gophercises "./gophercises"
)

func main() {

	fmt.Printf("########## Starting Exercise ##########\n\n")

	// 	Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

	// Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should
	//ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

	// Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one
	// at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

	// At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions
	// given invalid answers or unanswered are considered incorrect.

	// Bonus
	// As a bonus exercises you can also…

	// Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. Hint:
	// Check out the strings package.
	// Add an option (a new flag) to shuffle the quiz order each time it is run.

	gophercises.Quiz()

	fmt.Printf("\n########## Ending Exercise ##########\n")

}
