package main

import (
	//For Printing to the terminal

	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	fmt.Printf("########## Starting Exercise ##########\n\n")
	//Using the X axis to represent east-west direction and the Y axis to represent north-south direction
	//The rover starts by moving west, hence the direction coordinates are (1, 0)

	//Initialize Earth Satelites to receive messages
	fmt.Printf("\nmain(): Initialize Earth Satelites\n")
	msgToEarchChan := NewEarthMessager()
	time.Sleep(3 * time.Second)

	// Initialize a Grid on Mars
	fmt.Printf("\nmain(): Initialize a new Grid on Mars\n")
	mg := NewMarsGrid()
	time.Sleep(1 * time.Second)
	//mg.Print()

	//Create multiple rovers
	numOfRovers := 5

	//fmt.Printf("\nmain(): Initialize %v new Rovers on Mars\n", numOfRovers)

	rovers := make([]*Rover, numOfRovers)
	for i := range rovers {
		tmpName := "r" + strconv.Itoa(i)
		fmt.Printf("\nmain(): Initialise new Rover %v on Mars\n", tmpName)
		tmpRover := NewRover(mg, msgToEarchChan, tmpName)
		rovers[i] = tmpRover
		time.Sleep(3 * time.Second)
	}

	//Move the rovers around randomely
	for _, rover := range rovers {
		if rand.Intn(2) == 0 {
			fmt.Printf("\nmain(): Ask the rover %v to turn Right\n", rover.name)
			rover.Right()
		} else {
			fmt.Printf("\nmain(): Ask the rover %v to turn Left\n", rover.name)
			rover.Left()
		}
		time.Sleep(4 * time.Second)
	}

	//Stop the Rovers?

	time.Sleep(10 * time.Second)
	mg.Print()

	fmt.Printf("\n########## Ending Exercise ##########\n")

}
