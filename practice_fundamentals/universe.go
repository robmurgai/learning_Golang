// Exercise 20
// For this challenge, you will build a simulation of underpopulation, overpopulation, and reproduction called Conway’s Game of Life (see mng.bz/xOyY).
// The simulation is played out on a two-dimensional grid of cells. As such, this challenge focuses on slices.
//
// Each cell has eight adjacent cells in the horizontal, vertical, and diagonal directions. In each generation, cells live or die based on the number
// of living neighbors.

package practicefundamentals

import (
	"fmt"
	"math/rand"
	"time"
)

// For your first implementation of the Game of Life, limit the universe to a fixed size. Decide on the dimensions of the grid and define some constants
const (
	universeWidth  = 80 //columns
	universeHeight = 15 //rows
	alive          = "*"
	dead           = " "
)

// Universe type to hold a two-dimensional field of cells. With a Boolean type, each cell will be either dead (false) or alive (true)
// Uses slices rather than arrays so that a universe can be shared with, and modified by, functions or methods.
type Universe [][]bool

// NewUniverse uses make to allocate and return a Universe with height rows and width columns per row
// Freshly allocated slices will default to the zero value, which is false, so the universe begins empty.
func NewUniverse() Universe {

	//initialize univesre as type Universe and length universeHeight
	universe := make(Universe, universeHeight)
	//fmt.Printf("\nNewUniverse(): universe is decalared with height: %v\n", len(universe))
	//fmt.Println(universe)
	for i := range universe {

		//initialize a univesre[i] as type []bool and length universeWidth
		universe[i] = make([]bool, universeWidth)

	}
	//fmt.Printf("NewUniverse(): universe has been updated with height: %v and width: %v\n", len(universe), len(universe[0]))
	return universe
}

// Show prints a universe to the screen using the fmt package. It represents live cells with an asterisk and dead cells with a space.
// Be sure to move to a new line after printing each row.
// Also be sure that you can run your program, even though the universe is empty.
func (u Universe) Show() {
	for i := range u {
		fmt.Printf("Row %2v: ", i)
		for j := range u[i] {
			if u[i][j] {
				fmt.Printf(alive)
			} else {
				fmt.Printf(dead)
			}
			time.Sleep(1000000)
		}
		fmt.Println()
	}
	fmt.Println()

}

// Seed method that randomly sets approximately 25% of the cells to alive (true):
// Remember to import math/rand to use the Intn function. When you’re done, update main to populate the universe with Seed and display your handiwork
// with Show.
func (u Universe) Seed() {
	//fmt.Printf("\nSeed(): \n")

	numOfCellsSeeded := 0
	for i := range u {
		for j := range u[i] {
			if rand.Intn(4) == 3 {
				u[i][j] = true
				numOfCellsSeeded++
			}
		}
	}
	//fmt.Printf("Seed(): FYI: This Universe has been seeded with %v out of %v cells\n", numOfCellsSeeded, len(u)*len(u[0]))
}

// Alive is a way to determine whether a cell is alive'// The rules of Conway’s Game of Life are as follows:
// To implement the rules, break them down into three steps, each of which can be a method:
func (u Universe) Alive(row int, column int) bool {
	//fmt.Printf("\nAlive(): Universe u[%v][%v] is Alive: %v\n", row, column, u[row][column])
	return u[row][column]
}

// Neighbors is a method to count the number of live neighbors for a given cell, from 0 to 8. Rather than access the universe data directly, use the
// Alive method so that the universe wraps around.
// A complication arises when the cell is outside of the universe. Is (–1,–1) dead or alive? On an 80 × 15 grid, is (80,15) dead or alive?
// To address this, make the universe wrap around. The neighbor above (0,0) will be (0,14) instead of (0,–1), which can be calculated by adding height
// to y. If y exceeds the height of the grid, you can turn to the modulus operator (%) that we used for leap year calculations. Use % to divide y by
// height and keep the remainder. The same goes for x and width.
func (u Universe) Neighbors(row int, column int) int {

	aliveNeighbors := 0
	//(row-1, column-1)(row-1, column)(row-1, column+1)
	//(row, column-11)(row, column)(row, column+1)
	//(row+1, column-1)(row+1, column)(row+1, column+1)

	//fmt.Printf("\nNeighbors(): Received request to find alive neighbors for Universe u[%v][%v]\n", row, column)

	//fmt.Printf("\nNeighbors(): Universe u[%v][%v] is Alive: %v ", row, column, u[row][column])

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checkRow := (row - 1 + i + universeHeight) % (universeHeight)
			checkColumn := (column - 1 + j + universeWidth) % (universeWidth)
			//fmt.Printf("\nNeighbors(): Requesting to see if Neighbour u[%v][%v] is Alive\n", checkRow, checkColumn)
			if u.Alive(checkRow, checkColumn) {
				aliveNeighbors++
				//fmt.Printf("Neighbors(): Universe u[%v][%v] is (alive): %v.  Alive Neighbors is now: %v\n", checkRow, checkColumn, u[checkRow][checkColumn], aliveNeighbors)
			}
		}
	}
	//fmt.Printf("and it has %v alive neighbors\n", aliveNeighbors-1)
	return aliveNeighbors - 1 //in case the main cell u[row][height] was alive, we have to not count it, this is the 8 neighbors
}

// Next determines whether a cell should be alive or dead in the next generation
// Now that you can determine whether a cell has two, three, or more neighbors, you can implement the rules shown at the beginning of this section.
// Write a Next method to do this:
// A live cell with less than two live neighbors dies.
// A live cell with two or three live neighbors lives on to the next generation.
// A live cell with more than three live neighbors dies.
// A dead cell with exactly three live neighbors becomes a live cell.
func (u Universe) Next(row, column int) bool {

	//fmt.Printf("Next(): Calling u.Neighbours(%v, %v)\n", row, column)
	aliveNeighbors := u.Neighbors(row, column)
	//fmt.Printf("\nNext(): Universe u[%v][%v] with %v alive neighbors", row, column, aliveNeighbors)

	switch {
	case (u.Alive(row, column) && (aliveNeighbors == 2 || aliveNeighbors == 3)):
		//fmt.Printf(" will live with two or three live neighbors lives on to the next generation.")
		return true
	case (!u.Alive(row, column) && (aliveNeighbors == 3)):
		//fmt.Printf(" with exactly three live neighbors lives on to the next generation.")
		return true
	default:
		//fmt.Printf(" will die in the next generation")
		return false
	}
}

// Step through each cell in the universe and determine what its Next state should be.
// There’s one catch. When counting neighbors, your count should be based on the previous state of the universe. If you modify the universe directly,
// those changes will influence the neighbor counts for the surrounding cells.
// A simple solution is to create two universes of the same size. Read through universe A while setting cells in universe B. Write a Step function to
// perform this operation:
func Step(currentGen, nextGen Universe) {

	//fmt.Printf("\nStep() through the Universe\n")

	for row := 0; row < universeHeight; row++ {
		for column := 0; column < universeWidth; column++ {
			//fmt.Printf("Step() calling u.Next(%v, %v)\n", row, column)
			nextGen[row][column] = currentGen.Next(row, column)
			if nextGen[row][column] != currentGen[row][column] {
				//fmt.Printf("Step(): items in Row: %v and Column: %v will change in the next Gen\n", row, column)
			}
		}
	}
}
