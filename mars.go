package main

import (
	"fmt"
	"math/rand"
)

/**
**
** Lesson 5. Capstone: Ticket to Mars
** Welcome to the first challenge. It’s time to take everything covered in unit 1 and write a program on your own. Your challenge is
** to write a ticket generator in the Go Playground that makes use of variables, constants, switch, if, and for. It should also draw on
** the fmt and math/rand packages to display and align text and to generate random numbers.
**
** When planning a trip to Mars, it would be handy to have ticket pricing from multiple spacelines in one place. Websites exist that
** aggregate ticket prices for airlines, but so far nothing exists for spacelines. That’s not a problem for you, though. You can use
** Go to teach your computer to solve problems like this.
**
**
**
** Start by building a prototype that generates 10 random tickets and displays them in a tabular format with a nice header, as follows:
**
** Spaceline        Days Trip type  Price
** ======================================
** Virgin Galactic    23 Round-trip $  96
** Virgin Galactic    39 One-way    $  37
** SpaceX             31 One-way    $  41
** Space Adventures   22 Round-trip $ 100
** Space Adventures   22 One-way    $  50
** Virgin Galactic    30 Round-trip $  84
** Virgin Galactic    24 Round-trip $  94
** Space Adventures   27 One-way    $  44
** Space Adventures   28 Round-trip $  86
** SpaceX             41 Round-trip $  72
** The table should have four columns:
**
** The spaceline company providing the service
** The duration in days for the trip to Mars (one-way)
** Whether the price covers a return trip
** The price in millions of dollars
** For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.
**
** Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km away from Earth at the time.
**
** Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine the duration for the trip to Mars and also
** the ticket price. Make faster ships more expensive, ranging in price from $36 million to $50 million. Double the price for round
** trips.
**
** When you’re done, post your solution to the Get Programming with Go forums at forums.manning.com/forums/get-programming-with-go. If
** you get stuck, feel free to ask questions on the forums, or take a peek at the appendix for our solution.
**
**/

func lesson500() {

	const travelDistance = 62100000 // KM

	fmt.Printf("Spaceline        Days Trip type    Price\n")
	fmt.Printf("========================================\n")

	for count := 0; count < 10; count++ {

		travelSpeed := rand.Intn(15) + 16
		travelDays := ((travelDistance / travelSpeed) / 360) / 24
		//fmt.Printf("%v km/s %v days\n", travelSpeed, travelDays)
		travelPrice := travelSpeed + 20

		switch rand.Intn(3) {
		case 0:
			fmt.Printf("%-18v", "Virgin Galactic")
		case 1:
			fmt.Printf("%-18v", "SpaceX")
		default:
			fmt.Printf("%-18v", "Space Adventures")
		}

		fmt.Printf("%-4v", travelDays)

		switch rand.Intn(2) {
		case 0:
			fmt.Printf("%-11v $  %2v\n", "Round-trip", 2*travelPrice)
		case 1:
			fmt.Printf("%-11v $  %2v\n", "one-way", travelPrice)
		}
	}
}
