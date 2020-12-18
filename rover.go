// rover.go
// Using the following RoverDriver type as a starting point, define Start and Stop methods and associated commands and make the rover obey them.

package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	stop  = command(2)
)

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc  chan command
	pos       image.Point
	direction image.Point
}

// NewRoverDriver Constructor
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc:  make(chan command),
		pos:       image.Point{X: 0, Y: 0},
		direction: image.Point{X: 1, Y: 0},
	}
	return r
}

// drive is responsible for printing the current position of the rover. It
// is expected to be started in a goroutine.
func (r *RoverDriver) logPosition() {
	updateInterval := 250 * time.Millisecond
	for {
		time.Sleep(updateInterval)
		r.pos = r.pos.Add(r.direction)
		log.Printf("moved to %v", r.pos)
	}

}

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine.
func (r *RoverDriver) drive() {
	for {
		if c, ok := <-r.commandc; ok {
			if c == right {
				r.direction = image.Point{
					X: -r.direction.Y,
					Y: r.direction.X,
				}
			} else if c == left {
				r.direction = image.Point{
					X: r.direction.Y,
					Y: -r.direction.X,
				}
			}
			log.Printf("new direction %v", r.direction)
		} else {
			return
		}
	}
}

// Left turns the rover left (90° counterclockwise).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right (90° clockwise).
func (r *RoverDriver) Right() {
	r.commandc <- right
}

// Start the rover.
func (r *RoverDriver) Start() {
	log.Printf("Starting the rover at %v \n", r.pos)
	go r.drive()
	go r.logPosition()
}

// Stop the rover
func (r *RoverDriver) Stop() {
	log.Printf("Stopping at %v", r.pos)
	close(r.commandc)
}
