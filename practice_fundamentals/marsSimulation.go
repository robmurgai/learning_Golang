package practicefundamentals

import (
	"fmt"
	"math/rand"
	"time"
)

//  Make a few types of animals. Each animal should have a name
type animals struct {
	name string
}

//	Each animal should adhere to the Stringer
//	interface to return their name. For reference the fmt package declares a Stringer interface as follows:
//  type Stringer interface {
//     String() string
//  }
//  If a type provides a String method, Println, Sprintf, and friends will use it.
func (an animals) String() string {
	return fmt.Sprintf("Ark: We have %v on MARS! Huzzah! ", an.name)
}

// Every animal should have methods to move and eat. The move method should return a description of the movement. The eat method should return the name of a random food
// that the animal likes.
func (an animals) move() string {
	return fmt.Sprintf("Ark: %v has moved 10 paces.", an.name)
}

func (an animals) eat() string {
	return fmt.Sprintf("Ark: %v just ate some food!", an.name)
}

//  Time on Mars Struct
type marsTime struct {
	sol  int
	hour int
	min  int
}

// Implement a day/night cycle
const (
	solhours = 24
	sunrise  = 6  //Sunrise at 6 AM
	sunset   = 18 //Sunset at 6 PM
)

type dayNightCycler interface {
	isItDayTime() bool
}

func (mt marsTime) isItDayTime() bool {
	if mt.hour >= sunrise && mt.hour < sunset {
		return true
	}
	fmt.Println("Ark: Night time on Mars")
	return false
}

//	marsSimulation runs the simulation for three 24-hour sols (72 hours).
// 	All the animals should sleep from sunset until sunrise.
//	For every hour of the day,
// 	pick an animal at random to perform a random action (move or eat). For every action, print out a description of what the animal did.
func marsSimulation() {

	// Start the Simulation :)
	fmt.Println("Ark: Starting Simulation")

	// Current Time Set to Day 1, time 00:00
	currentMT := marsTime{hour: 0, min: 0, sol: 1}
	fmt.Printf("Ark: Current Time on Mars: %+v\n", currentMT)

	//Create the Ark.
	marsArk := []animals{
		{name: "girrafs"},
		{name: "deer"},
		{name: "monkey"},
		{name: "dogs"},
		{name: "rabbits"},
	}

	fmt.Printf("Ark: The Mars Ark has %v animals\n", len(marsArk))
	for _, animal := range marsArk {
		fmt.Println(animal)
	}

	//Start the 3 day simluation
	for currentMT.sol < 4 {

		fmt.Printf("Ark: Current time on Mars is %+v\n", currentMT)

		if currentMT.isItDayTime() {

			//Lets randomaly select an animal from the ark
			randomAnimal := marsArk[rand.Intn(len(marsArk))]

			//Lets randomly select to eat or Move
			if rand.Intn(2) == 0 {
				fmt.Println(randomAnimal.move())
			} else {
				fmt.Println(randomAnimal.eat())
			}
		}

		//increment current time by 1 hour
		currentMT.hour++
		if currentMT.hour > solhours {
			//increment current time by 1 day
			currentMT.sol++
			currentMT.hour %= solhours
		}
		time.Sleep(2 * time.Second)

	}

	// End the Simulation :)
	fmt.Println("Ark: Ending Simulation")
}
