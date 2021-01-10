package gophercises

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Exercise details
// This exercise is broken into two parts to help simplify the process of explaining it as well as to make it easier to solve. The second part is
// harder than the first, so if you get stuck feel free to move on to another problem then come back to part 2 later.
//
// Part 1
// Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how
// many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be
// asked immediately afterwards.
//
// The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.
//
// The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that
// question.

// 5+5,10
// 7+3,10
// 1+1,2
// 8+3,11
// 1+2,3
// 8+6,14
// 3+1,4
// 1+4,5
// 5+1,6
// 2+3,5
// 3+3,6
// 2+4,6
// 5+2,7
// You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

// At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given
// invalid answers are considered incorrect.

// NOTE: CSV files may have questions with commas in them. Eg: "what 2+2, sir?",4 is a valid row in a CSV. I suggest you look into the CSV package in
// Go and don’t try to write your own CSV parser.

// Part 2
// Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.
//
// Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn’t wait for the user to answer one final questions but should
// ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

// Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one
// at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

// At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions
// given invalid answers or unanswered are considered incorrect.

// Bonus:  As a bonus exercises you can also…
// 		Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect.
//		Check out the strings package.
// 		Add an option (a new flag) to shuffle the quiz order each time it is run.

// Program Flow
// Run the program with arugment for file name.
// 		store file name if available.
//		Use file name is not
// Open a quiz file
// Default to problems.csv
// Read each line and print to screen
// store the asnwer and print next line.
// When done, print the score. (X correct out of Y)

type quizFileLine struct {
	qns string
	key string
	ans string
}

// Quiz will read in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how
// many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be
// asked immediately afterwards.
func Quiz() {

	fmt.Println("Start Quiz")

	// Quiz File Name
	var quizFileName string
	flag.StringVar(&quizFileName, "csv", "problems.csv", "<quiz_file_name>")

	// Quiz Timer
	timerPtr := flag.Int("t", 30, "<quiz timer (seconds)>")

	// Quiz File Shuffle
	var quizFileShuffle bool
	flag.BoolVar(&quizFileShuffle, "s", false, "Shuffle the quiz")

	flag.Parse()
	// log.Printf("DEBUG: quiz(): quizFileName: %v\n", quizFileName)
	// log.Printf("DEBUG: quiz(): timer: %v\n", *timerPtr)
	// log.Printf("DEBUG: quiz(): quizFileShuffle: %v\n", quizFileShuffle)

	// Warn about arguments we are ignoring
	if len(flag.Args()) != 0 {
		log.Printf("WARN: quiz(): Ignoring the following arguments: %v\n", flag.Args())
	}

	//Env independent quizfile path
	dir1 := "gophercises"
	quizFilePath := filepath.Join(dir1, quizFileName)
	//log.Printf("DEBUG: quiz(): quizFilePath: %v\n", quizFilePath)

	// Open and read the quiz file
	quizFile, err := os.Open(quizFilePath)
	if err != nil {
		log.Printf("ERROR: quiz(): Unable to Open file: %v, %v\n", quizFilePath, err)
		panic(err)
	}
	defer quizFile.Close()

	// Read the CSV Quiz file
	r := csv.NewReader(quizFile)

	fullFile, err := r.ReadAll()
	if err != nil {
		log.Printf("ERROR: quiz(): Unable to Read and Parse CSV file: %v", quizFilePath)
		panic(err)
	}
	problemTotal := len(fullFile)
	//log.Printf("DEBUG: quiz(): problemTotal: %d\n", problemTotal)

	if quizFileShuffle {
		//log.Printf("DEBUG: quiz(): Before shuffling fullFile: %v\n", fullFile)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(fullFile), func(i, j int) { fullFile[i], fullFile[j] = fullFile[j], fullFile[i] })
		//log.Printf("DEBUG: quiz(): After shuffling fullFile: %v\n", fullFile)
	}

	//Start the timer
	fmt.Printf("You have %d seconds to answer %d questions, press enter when ready: ", *timerPtr, problemTotal)
	var startPrompt string
	fmt.Scanln(&startPrompt)
	fmt.Println()

	//start the timer;
	timeout := time.After(time.Duration(*timerPtr) * time.Second)

	//start the quiz
	var score int
	qfls := make([]quizFileLine, problemTotal)
	quizChan := make(chan bool)
	go startQuiz(fullFile, &score, qfls, quizChan)

	//when timer expires or the quiz completes, end the quiz
	select {
	case <-quizChan:
		//log.Printf("\nDEBUG: quiz(): Quiz Complete. \n%v \nAnswered %v questions correctly out of %v\n", qfls, score, problemTotal)
		fmt.Println()
		endQuiz(score, problemTotal)
	case <-timeout:
		//log.Printf("\nDEBUG: quiz(): Timeout. \n%v \nAnswered %v questions correctly out of %v\n", qfls, score, problemTotal)
		fmt.Println()
		fmt.Printf("\nTimesUp.  ")
		endQuiz(score, problemTotal)
	}

}

func startQuiz(fullFile [][]string, scorePtr *int, qfls []quizFileLine, quizChan chan bool) {

	for problemNum, record := range fullFile {

		//log.Printf("DEBUG: quiz(): Record: %v\n", record)
		if len(record) != 2 {
			errString := fmt.Sprintf("Expected 2 comma seperated values per record in the CSV file: [Question Answer-Key] \n") +
				fmt.Sprintf("Record[%v] has %v comma seperated values: %v \n", problemNum+1, len(record), record)
			//log.Printf("ERROR: quiz(): %v", errString)
			panic(errString)
		}

		qns := record[0]
		keyInt, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("ERROR: quiz(): Unable to convert key %v from string to integer %v\n", record[1], err)
			panic(err)
		}
		//log.Printf("DEBUG: quiz(): key: %d\n", key)
		fmt.Printf("Problem #%v: %v = ", problemNum+1, qns)

		var ans string
		fmt.Scanln(&ans)

		//log.Printf("DEBUG: quiz(): ans: %s\n", ans)
		if ans != "" {
			ansInt, err := strconv.Atoi(ans)
			if err != nil {
				log.Printf("WARN: quiz(): Record[%v]: %v, unable to convert user response: %q to integer value.  %v\n", problemNum, record, ans, err)
			} else if keyInt == ansInt {
				*scorePtr++
				//log.Printf("DEBUG: quiz(): Answerred Correctly\n")
			}
		}

		qfls[problemNum] = quizFileLine{qns, record[1], ans}
		problemNum++
	}
	//log.Printf("DEBUG: startQuiz(): All Queations Answered: \n%v\n", qfls)
	quizChan <- true
	return
}

func endQuiz(score int, problemTotal int) {
	fmt.Println("Quiz Complete")
	fmt.Printf("\nYou scored %v out of %v\n", score, problemTotal)
}
