// Practice using goroutines and mutexes
// A GRID TO ROVE ON
// Make a grid that the rover can drive around on by implementing a MarsGrid type. You’ll need to use a mutex to make it safe for use by multiple
// goroutines at once.
//
// We want to find life on Mars, so we’ll send several rovers down to search for it, but we need to know when life is found. In every cell in
// the grid, assign some likelihood of life, a random number between 0 and 1000.

package main

import (
	"fmt"
	"image"
	"math/rand"
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
	gridLife   map[image.Point]int
}

// NewMarsGrid constructs a new MarsGrid
func NewMarsGrid() *MarsGrid {
	var mg MarsGrid
	mg.grid = make(map[image.Point]*Occupier, 2*gridLength+1)
	mg.gridLife = make(map[image.Point]int, 2*gridLength+1)
	for i := -gridLength; i <= gridLength; i++ {
		for j := -gridLength; j <= gridLength; j++ {
			tmpPoint := image.Point{X: i, Y: j}
			mg.grid[tmpPoint] = nil
			mg.gridLife[tmpPoint] = rand.Intn(1000)
		}
	}
	mg.parkingLot = marsParkingLot

	//Setting life at Parking lot to 0, to keep things simple.
	mg.gridLife[marsParkingLot] = 0
	// log.Printf("MarsGrid.NewMarsGrid(): DEBUG New Mars grid created\n")
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

	// log.Printf("MarsGrid.Occupy(): DEBUG: Request received to initialize a new occupier at location: %v\n", p)

	// Lock to enable one read/write of the grid
	// log.Printf("DEBUG: MarsGrid.Occupy(): Mars grid lock requested")
	// g.lock()
	// log.Printf("DEBUG: MarsGrid.Occupy(): Mars grid unlock requested")
	// defer g.unlock()
	if !g.locationAvailable(p) {
		return nil
	}
	var o Occupier
	o.grid = g
	o.pos = p
	g.grid[p] = &o

	// log.Printf("DEBUG: MarsGrid.Occupy(): Request Successful. A new Occupier %v initialized at location %v\n", o, p)
	return &o

}

// Move moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside
// the grid or because the cell it's trying to move into
// is occupied. If it fails, the occupier remains in the same place.
func (ocp *Occupier) Move(p image.Point) bool {
	// log.Printf("DEBUG: MarsGrid.Move(): %v: Request received to move to location: %v\n", ocp.name, p)

	g := ocp.grid

	// Lock the grid to enable one read/write while trying to move, to avoid
	//log.Printf("DEBUG: MarsGrid.Move(): Lock Requested by %v: to enable move to location %v", ocp.name, p)
	g.lock(ocp.name, p)
	//log.Printf("DEBUG: MarsGrid.Move(): Unlock Scheduled by %v: to enable move to location %v", ocp.name, p)
	defer g.unlock(ocp.name, p)

	if !g.locationAvailable(p) {
		return false
	}

	//Set current location to nil
	g.grid[ocp.pos] = nil

	//Update Occupier to this location
	ocp.pos = p

	//Update new location on grid with the occupier.
	g.grid[p] = ocp

	// log.Printf("DEBUG: MarsGrid.Move(): %v: Request Successful. Moved to location %v\n", ocp.name, p)
	return true

}

// Print the grid
func (g *MarsGrid) Print() {
	for i := -gridLength; i <= gridLength; i++ {
		for j := -gridLength; j <= gridLength; j++ {
			printVal := fmt.Sprintf("%-4s", "_")
			tmpMapKey := image.Point{X: i, Y: j}
			if tmpMapKey == marsParkingLot {
				printVal = fmt.Sprintf("%-4s", "P")
			}

			var tmpPrintVal string
			if g.likelihoodOfLife(tmpMapKey) {
				tmpPrintVal = "*"
				printVal = fmt.Sprintf("%-4s", tmpPrintVal)
			}
			if occupier, ok := g.grid[tmpMapKey]; ok {
				if occupier != nil {
					tmpPrintVal += occupier.name
					printVal = fmt.Sprintf("%-4s", tmpPrintVal)
				}
			}
			fmt.Printf("%v", printVal)
		}
		fmt.Println()
	}
}

func (g *MarsGrid) locationAvailable(p image.Point) bool {
	tmpOccupier, ok := g.grid[p]
	if !ok { // outside the grid
		// log.Printf("DEBUG: MarsGrid.locationNotAvailable(): %v is outside the grid\n", p)
		return false
	} else if tmpOccupier != nil { //occupied
		if p != marsParkingLot {
			// log.Printf("DEBUG: MarsGrid.locationNotAvailable(): %v: is already on location %v.  Request Denied.\n", tmpOccupier.name, p)
			return false
		}
	}
	return true
}

// If a rover finds a cell with a life value above 900, it may have
// found life and it must send a radio message back to Earth.
func (g *MarsGrid) likelihoodOfLife(p image.Point) bool {
	if g.gridLife[p] > 900 {
		return true
	}
	return false
}

//Lock and Unlockfunctions for debugging
func (g *MarsGrid) lock(ocpName string, p image.Point) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("ERROR: MarsGrid.lock(): When requested by %v: to enable move to location %v\n", ocpName, p)
			fmt.Println("ERROR: MarsGrid.lock(): ", r)
			panic(r)
		}
	}()
	//log.Printf("DEBUG: MarsGrid.lock(): Requested by %v: to enable move to location %v", ocpName, p)
	g.mu.Lock()
	//log.Printf("DEBUG: MarsGrid.lock(): Successful by %v: to enable move to location %v", ocpName, p)
	return true
}

func (g *MarsGrid) unlock(ocpName string, p image.Point) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("ERROR: MarsGrid.unlock(): When requested by %v: to enable move to location %v\n", ocpName, p)
			fmt.Println("ERROR: MarsGrid.unlock(): ", r)
			panic(r)
		}
	}()
	//log.Printf("DEBUG: MarsGrid.unlock(): Requested by %v: to enable move to location %v", ocpName, p)
	g.mu.Unlock()
	//log.Printf("DEBUG: MarsGrid.unlock(): Successful by %v: to enable move to location %v", ocpName, p)
	return true
}
