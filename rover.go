// Practice using goroutines and mutexes
// Using the following RoverDriver type as a starting point, define Start and Stop methods and associated commands and make the rover obey them.
// Have the rover use a MarsGrid object passed into the NewRover function. If it hits the edge of the grid or an obstacle, it should turn and
// go in another random direction.
//
// ImagelPOints are X, Y coordinates which increase to the right and down (not up)

package main

import (
	"image"
	"log"
	"math/rand"
	"time"
)

// Rover drives around the surface of a planet based on the direction command it receives.
// It essentially only needs a channel to receive commands
type Rover struct {
	commandChan  chan command
	gridOccupier *Occupier
	name         string
}

// NewRover Constructor: Initialize a new Rover with a command channel and start driving by calling drive()
func NewRover(grid *MarsGrid, name string) *Rover {
	r := Rover{
		commandChan: make(chan command),
		name:        name,
	}

	//Set Rover as an occupier in the Parking Lot on a planet.
	r.gridOccupier = grid.Occupy(grid.parkingLot)
	if r.gridOccupier == nil {
		log.Panicf("ERROR: Unable to Initialize a new Rover on this planet")
		return nil
	}

	r.gridOccupier.name = r.name
	log.Printf("NewRover(): New Rover %v Initialized on Mars.  Current Location is the Rover Parking Lot %v\n", r.name, grid.parkingLot)

	go r.drive()
	return &r
}

// RoverParking is a Parking Lot for new rovers where its okay to have more than one rover.

// drive is responsible for driving the rover and turning it left or right, depending on the direction command
// received.  The rover moves forward at the speed of four steps a second (or 1 step every 250 milli seconds)
func (r *Rover) drive() {

	log.Printf("drive(): Rover %v is moving\n", r.name)

	// Current Direction - Moving East
	direction := image.Point{X: 1, Y: 0}
	log.Printf("drive(): %v Current direction -> right %v", r.name, direction)

	// Speed is 1 step per 250 seconds
	timePerStep := 250 * time.Millisecond
	timeToMoveForwardOneStep := time.After(timePerStep)
	for {
		select {
		case c := <-r.commandChan: // If we receive a command to change direction
			switch c {
			case right: // Its a command to turn right
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
				log.Printf("drive(): %v New direction -> right %v", r.name, direction)
			case left: // Its a command to turn left
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
				log.Printf("drive(): %v New direction -> left %v", r.name, direction)
			} // End of switch c
		case <-timeToMoveForwardOneStep:
			tmpCurrentPosition := r.gridOccupier.pos
			tmpNewPos := tmpCurrentPosition.Add(direction)
			log.Printf("drive(): DEBUG %v Time to move forward one step\n", r.name)
			log.Printf("drive(): DEBUG %v Current Position: %v\n", r.name, tmpCurrentPosition)
			log.Printf("drive(): DEBUG %v New Position Requested: %v\n", r.name, tmpNewPos)
			tmpD := direction
			for !r.gridOccupier.Move(tmpNewPos) {
				log.Printf("drive(): %v %v grid location not available, picking a new direction", r.name, tmpNewPos)
				switch command(rand.Intn(2)) {
				case right: // Its a command to turn right
					tmpD = image.Point{
						X: -direction.Y,
						Y: direction.X,
					}
					log.Printf("drive(): %v New direction -> right %v", r.name, tmpD)
					tmpNewPos = tmpCurrentPosition.Add(tmpD)
				case left: // Its a command to turn left
					tmpD = image.Point{
						X: direction.Y,
						Y: -direction.X,
					}
					log.Printf("drive(): %v New direction -> left %v", r.name, tmpD)
					tmpNewPos = tmpCurrentPosition.Add(tmpD)
				} // End of switch c
			} // End of for loop
			log.Printf("drive(): %v Rover moved to %v", r.name, r.gridOccupier.pos)
			direction = tmpD
			timeToMoveForwardOneStep = time.After(timePerStep)
		} // End of Select
	} // End of for loop
}

// Left turns the rover left (90° counter-clockwise).
func (r *Rover) Left() {
	r.commandChan <- left
}

// Right turns the rover right (90° clockwise).
func (r *Rover) Right() {
	r.commandChan <- right
}

// Commands come in the form of right or left.
type command int

const (
	right = command(0)
	left  = command(1)
)
