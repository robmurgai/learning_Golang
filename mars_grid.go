// Practice using goroutines and mutexes
// A GRID TO ROVE ON
// Make a grid that the rover can drive around on by implementing a MarsGrid type. You’ll need to use a mutex to make it safe for use by multiple
// goroutines at once. It should look something like the following:

package main

import (
	"fmt"
	"image"
	"log"
	"sync"
)

const gridLength = 10

var marsParkingLot = image.Point{X: 0, Y: 0}

// MarsGrid represents a grid of some of the surface
// of Mars. It may be used concurrently by different
// goroutines.
type MarsGrid struct {
	grid       map[image.Point]*Occupier
	parkingLot image.Point
	mu         sync.Mutex
}

// NewMarsGrid constructs a new MarsGrid
func NewMarsGrid() *MarsGrid {
	var mg MarsGrid
	mg.grid = make(map[image.Point]*Occupier, 2*gridLength+1)
	for i := -gridLength; i <= gridLength; i++ {
		for j := -gridLength; j <= gridLength; j++ {
			tmpPoint := image.Point{X: i, Y: j}
			mg.grid[tmpPoint] = nil
		}
	}
	mg.parkingLot = marsParkingLot
	log.Printf("MarsGrid.NewMarsGrid(): New Mars grid created\n")
	return &mg
}

// Occupier represents an object that can occupy a cell in the grid.  It may be a rover or something else
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
	name string
}

// Occupy occupies a cell at the given point in the grid. It
// returns nil if the point is already occupied or the point is
// outside the grid. Otherwise it returns a new occupier that can
// move to different places on the grid.
// Occupy is safe to use in multiple goroutines
func (g *MarsGrid) Occupy(p image.Point) *Occupier {

	log.Printf("MarsGrid.Occupy(): DEBUG: Request received to initialize a new occupier at location: %v\n", p)

	// Lock to enable one read/write of the grid
	g.mu.Lock()
	defer g.mu.Unlock()
	occupied, ok := g.grid[p]
	if !ok { // outside the grid
		log.Printf("MarsGrid.Occupy(): DEBUG: %v is outside the grid\n", p)
		return nil
	} else if occupied != nil { //occupied
		if p != marsParkingLot {
			log.Printf("MarsGrid.Occupy(): DEBUG: %v is already has an occupier on the grid\n", p)
			return nil
		}
	}
	var o Occupier
	o.grid = g
	o.pos = p

	g.grid[p] = &o

	log.Printf("MarsGrid.Occupy(): DEBUG: Request Successful\n")
	log.Printf("MarsGrid.Occupy(): DEBUG: A new Occupier %v initialized at location %v\n", o, p)

	return &o

}

// Move moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside
// the grid or because the cell it's trying to move into
// is occupied. If it fails, the occupier remains in the same place.
func (ocp *Occupier) Move(p image.Point) bool {
	log.Printf("MarsGrid.Move(): DEBUG: Request received to move the occupier to location: %v\n", p)

	mg := ocp.grid

	// Lock the grid to enable one read/write while trying to move, to avoid
	mg.mu.Lock()
	defer mg.mu.Unlock()

	tmpOccupier, ok := mg.grid[p]
	if !ok { // outside the grid
		log.Printf("MarsGrid.Move(): DEBUG: %v is outside the grid\n", p)
		return false
	} else if tmpOccupier != nil { //occupied
		if p != marsParkingLot {
			log.Printf("MarsGrid.Move(): DEBUG: %v is already occupied on the grid\n", p)
			return false
		}
	}

	//Set current location to nil
	mg.grid[ocp.pos] = nil

	//Update new location for Occupier.
	ocp.pos = p

	//Update new location on grid with the occupier.
	mg.grid[p] = ocp

	log.Printf("MarsGrid.Move(): DEBUG: Request Successful\n")
	return true

}

// Print the grid
func (g *MarsGrid) Print() {
	for i := -gridLength; i <= gridLength; i++ {
		for j := -gridLength; j <= gridLength; j++ {
			tmpMapKey := image.Point{X: i, Y: j}
			if occupier, ok := g.grid[tmpMapKey]; ok {
				if tmpMapKey == marsParkingLot {
					fmt.Printf("%-4s", "P")
				} else if occupier != nil {
					fmt.Printf("%-4s", occupier.name)
				} else {
					fmt.Printf("%-4s", "_")
				}
			}
		}
		fmt.Println()
	}
}
