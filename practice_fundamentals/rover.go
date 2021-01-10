// Practice using goroutines and mutexes
// Have the rover use a MarsGrid object passed into the NewRover function. If it hits the edge of the grid or an obstacle, it should turn and
// go in another random direction.
//
// ImagelPOints are X, Y coordinates which increase to the right and down (not up)
//
// We want to find life on Mars, so we’ll send several rovers down to search for it, but we need to know when life is found. In every cell in
// the grid, assign some likelihood of life, a random number between 0 and 1000. If a rover finds a cell with a life value above 900, it may have
// found life and it must send a radio message back to Earth.
//
// Unfortunately, it’s not always possible to send a message immediately because the relay satellite is not always above the horizon. Implement a
// buffer goroutine that receives messages sent from the rover and buffers them into a slice until they can be sent back to Earth.
//
// Implement Earth as a goroutine that receives messages only occasionally (in reality for a couple of hours every day, but you might want to make
// the interval a little shorter than that). Each message should contain the coordinates of the cell where the life might have been found, and the
// life value itself.
//
// You may also want to give a name to each of your rovers and include that in the message so you can see which rover sent it. It’s also helpful to
// include the name in the log messages printed by the rovers so you can track the progress of each one.
//
// Set your rovers free to search and see what they come up with!

package practicefundamentals

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"
)

// Rover drives around the surface of a planet based on the direction command it receives.
// It essentially only needs a channel to receive commands
type Rover struct {
	commandChan  chan command
	messageChan  chan string
	gridOccupier *Occupier
	name         string
}

// NewRover Constructor: Initialize a new Rover with a command channel and start driving by calling drive()
func NewRover(grid *MarsGrid, msgToEarthChan *chan string, name string) *Rover {
	r := Rover{
		commandChan: make(chan command),
		messageChan: *msgToEarthChan,
		name:        name,
	}

	//Set Rover as an occupier in the Parking Lot on a planet.
	r.gridOccupier = grid.Occupy(grid.parkingLot)
	if r.gridOccupier == nil {
		log.Panicf("ERROR: Unable to Initialize a new Rover on this planet")
		return nil
	}

	r.gridOccupier.name = r.name
	//log.Printf("DEBUG: NewRover(): New Rover %v Initialized on Mars.  Current Location is the Rover Parking Lot %v\n", r.name, grid.parkingLot)

	go r.drive()
	return &r
}

// RoverParking is a Parking Lot for new rovers where its okay to have more than one rover.

// drive is responsible for driving the rover and turning it left or right, depending on the direction command
// received.  The rover moves forward at the speed of four steps a second (or 1 step every 250 milli seconds)
func (r *Rover) drive() {

	//log.Printf("drive(): Rover %v is moving\n", r.name)

	// Current Position and direction (Moving East)
	pos := r.gridOccupier.pos
	direction := image.Point{X: 1, Y: 0}

	// Speed is 1 step per 250 seconds
	timePerStep := 250 * time.Millisecond
	timeToMoveForwardOneStep := time.After(timePerStep)

	for {
		select {
		case c := <-r.commandChan: // Received a command to change direction
			//log.Printf("DEBUG: drive(): %v: Change direction", r.name)
			direction = updateDirection(direction, c)
		case <-timeToMoveForwardOneStep: // Time to move forward one step
			tmpNewPos := pos.Add(direction)
			tmpD := direction

			//Keep trying new direction till rover is able to move forward.
			for !r.gridOccupier.Move(tmpNewPos) {
				//log.Printf("DEBUG: drive(): %v: Unable to move forward to: %v", r.name, tmpNewPos)
				tmpD = updateDirection(direction, command(rand.Intn(2)))
				tmpNewPos = pos.Add(tmpD)
				//log.Printf("DEBUG drive(): %v: trying new direction to: %v", r.name, tmpNewPos)
			}
			// Rover successfull in move to new location
			pos = r.gridOccupier.pos
			direction = tmpD
			//log.Printf("DEBUG: drive(): %v: %v", r.name, pos)

			// Process Likelyhood of Life on this grid location
			r.foundLife()

			// Restart the timer to move forward one step.
			timeToMoveForwardOneStep = time.After(timePerStep)
		}
	}
}

// Left turns the rover left (90° counter-clockwise).
func (r *Rover) Left() {
	//log.Printf("DEBUG: Left(): %v: Sending Command to move Left\n", r.name)
	r.commandChan <- left
}

// Right turns the rover right (90° clockwise).
func (r *Rover) Right() {
	//log.Printf("DEBUG: Right(): %v: Sending Command to move Right\n", r.name)
	r.commandChan <- right
}

// Commands come in the form of right or left.
type command int

const (
	right = command(0)
	left  = command(1)
)

// foundLife checks rover's location for life.  If the location has a life value above 900, it send a radio message
// back to Earth.
func (r *Rover) foundLife() {
	g := r.gridOccupier.grid
	pos := r.gridOccupier.pos
	lifeValue := g.gridLife[pos]
	if g.likelihoodOfLife(pos) {
		msg := fmt.Sprintf("[%v(%v)%7v] Rover %[1]v found likelyhood of life (%[2]v) at location %[3]v\n", r.name, lifeValue, pos)
		//log.Printf("INFO: Rover.foundLife(): Sending message to Earth: %v", msg)
		r.messageChan <- msg
	}
}

func updateDirection(direction image.Point, c command) image.Point {
	switch c {
	case right: // Its a command to turn right
		direction = image.Point{
			X: -direction.Y,
			Y: direction.X,
		}
		// log.Printf("Right \n")
	case left: // Its a command to turn left
		direction = image.Point{
			X: direction.Y,
			Y: -direction.X,
		}
		// log.Printf("left \n")
	} // End of switch c
	return direction
}
