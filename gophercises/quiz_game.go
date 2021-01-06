package gophercises

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	key int
	ans int
}

// Quiz will read in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how
// many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be
// asked immediately afterwards.
func Quiz() {

	fmt.Println("Start Quiz")

	var quizFileName string
	flag.StringVar(&quizFileName, "csv", "problems.csv", "-csv <quiz_file_name>")

	flag.Parse()
	//log.Printf("DEBUG: quiz(): quizFileName: %v\n", quizFileName)

	// Warn about arguments we are ignoring
	if len(flag.Args()) != 0 {
		log.Printf("WARN: quiz(): Ignoring the following arguments: %v\n", flag.Args())
	}

	//Run the quiz file
	dir1 := "gophercises"
	quizFilePath := filepath.Join(dir1, quizFileName)
	//log.Printf("DEBUG: quiz(): quizFilePath: %v\n", quizFilePath)

	// Open the file and read one line at a time.
	quizFile, err := os.Open(quizFilePath)
	if err != nil {
		log.Printf("ERROR: quiz(): Unable to Open file: %v, %v\n", quizFilePath, err)
		panic(err)
	}
	defer quizFile.Close()

	r := csv.NewReader(quizFile)
	var score int
	var qfls []quizFileLine

	fullFile, err := r.ReadAll()
	if err != nil {
		log.Printf("ERROR: quiz(): Unable to Read file: %v, %v\n", quizFilePath, err)
		panic(err)
	}
	problemTotal := len(fullFile)
	//log.Printf("DEBUG: quiz(): problemTotal: %d\n", problemTotal)

	for problemNum, record := range fullFile {

		//log.Printf("DEBUG: quiz(): Record: %v\n", record)
		if len(record) != 2 {
			log.Printf("ERROR: quiz(): Unable to Read file: %v, %v\n", quizFilePath, err)
			panic(err)
		}
		qns := record[0]
		key, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("ERROR: quiz(): Unable to convert key %v from string to integer %v\n", record[1], err)
			panic(err)
		}
		fmt.Printf("Problem #%v: %v = ", problemNum+1, qns)

		var ans int
		fmt.Scanln(&ans)

		//log.Printf("DEBUG: quiz(): ans: %d\n", ans)
		//log.Printf("DEBUG: quiz(): key: %d\n", key)
		if ans == key {
			score++
			//log.Printf("DEBUG: quiz(): Answerred Correctly\n")
		}

		qfls = append(qfls, quizFileLine{qns, key, ans})
		problemNum++
	}

	defer displayScore(score, problemTotal)

}

func displayScore(score int, problemTotal int) {
	fmt.Printf("You scored %v out of %v", score, problemTotal)
}
