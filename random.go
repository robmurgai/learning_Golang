package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Generate a random year instead of always using 2018.
//For February, assign daysInMonth to 29 for leap years and 28 for other years.
//Hint: you can put an if statement inside of a case block.
//Use a for loop to generate and display 10 random dates.
func lesson414() {

	era := "AD"
	daysInMonth := 31
	currentYear := 2020

	for i := 11; i > 0; i-- {
		day := rand.Intn(daysInMonth) + 1
		year := rand.Intn(currentYear) + 1
		month := rand.Intn(12) + 1

		switch month {
		case 2:

			if (year % 4) == 0 {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		fmt.Println(era, year, month, day)
	}
}

// Write a guess-the-number program. Make the computer pick random numbers between 1–100 until it
// guesses your number, which you declare at the top of the program. Display each guess and whether it
// was too big or too small.
func lesson413() {
	var myNumber = 12
	var yourNumber = 0
	var youGotIt = false

	for !youGotIt {
		time.Sleep(time.Second)
		yourNumber = rand.Intn(100) + 1
		fmt.Printf("\nYou guessed %v", yourNumber)
		if yourNumber == myNumber {
			youGotIt = true
			fmt.Printf(" that was correct!")
		}

	}
}
