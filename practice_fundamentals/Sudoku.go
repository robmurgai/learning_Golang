// Practice Pointers, Errors, Nils,
// Sudoku is a logic puzzle that takes place on a 9 × 9 Grid (see en.wikipedia.org/wiki/Sudoku). Each square can contain a digit from 1 through 9.
// The number zero indicates an empty square.

// The Grid is divided into nine subregions that are 3 × 3 each. When placing a digit, it must adhere to certain constraints. The digit being placed
// may not already appear in any of the following:

// The horizontal row it’s placed in
// The vertical column it’s placed in
// The 3 × 3 subregion it’s placed in

// Use a fixed-size (9 × 9) array to hold the Sudoku Grid. If a function or method needs to modify the array, remember that you need to pass the
// array with a pointer.

// Implement a method to set a digit at a specific location. This method should return an error if placing the digit breaks one of the rules.

// Also implement a method to clear a digit from a square. This method need not adhere to these constraints, as several squares may be empty (zero).

// Sudoku puzzles begin with some digits already set. Write a constructor function to prepare the Sudoku puzzle, and use a composite literal to
// specify the initial values. Here’s an example:

// s := NewSudoku([rows][columns]int8{
//     {5, 3, 0, 0, 7, 0, 0, 0, 0},
//     {6, 0, 0, 1, 9, 5, 0, 0, 0},
//     {0, 9, 8, 0, 0, 0, 0, 6, 0},
//     {8, 0, 0, 0, 6, 0, 0, 0, 3},
//     {4, 0, 0, 8, 0, 3, 0, 0, 1},
//     {7, 0, 0, 0, 2, 0, 0, 0, 6},
//     {0, 6, 0, 0, 0, 0, 2, 8, 0},
//     {0, 0, 0, 4, 1, 9, 0, 0, 5},
//     {0, 0, 0, 0, 8, 0, 0, 7, 9},
// })
// The starting digits are fixed in place and may not be overwritten or cleared. Modify your program so that it can identify which digits are fixed
// and which are penciled in. Add a validation that causes set and clear to return an error for any of the fixed digits. The digits that are
// initially zero may be set, overwritten, and cleared.

// You don’t need to write a Sudoku solver for this exercise, but be sure to test that all the rules are implemented correctly.

// Use the following main funcation to see how the program works
// func main() {

// 	fmt.Printf("########## Starting Exercise ##########\n\n")

// 	s := NewSudokuGrid([9][9]int8{
// 		{5, 3, 0, 0, 7, 0, 0, 0, 0},
// 		{6, 0, 0, 1, 9, 5, 0, 0, 0},
// 		{0, 9, 8, 0, 0, 0, 0, 6, 0},
// 		{8, 0, 0, 0, 6, 0, 0, 0, 3},
// 		{4, 0, 0, 8, 0, 3, 0, 0, 1},
// 		{7, 0, 0, 0, 2, 0, 0, 0, 6},
// 		{0, 6, 0, 0, 0, 0, 2, 8, 0},
// 		{0, 0, 0, 4, 1, 9, 0, 0, 5},
// 		{0, 0, 0, 0, 8, 0, 0, 7, 9},
// 	})

// 	s.prettyPrint()

// 	for i := 0; i < 100; i++ {
// 		err := s.set(int8(rand.Intn(10)), int8(rand.Intn(9)), int8(rand.Intn(9)))
// 		if err != nil {
// 			fmt.Println(err)
// 			s.prettyPrint()
// 			fmt.Println()
// 		}
// 	}

// 	s.prettyPrint()
// 	fmt.Printf("\n########## Ending Exercise ##########\n")

// }

//When ready, expand the error with a SudokuError as a new error type
//Right now I don't now how to use this struct, so I am setting it up for v2 of this exercise.
// type SudokuError struct {
// 	LocationRow    int8
// 	LocationColumn int8
// 	LocationValue  int8
// 	Err            error
// }
// func (e *SudokuError) Error() string {
// 	return "Error: " + string(e.LocationValue) + " " + e.Err.Error() + " at (" + string(e.LocationRow) + "," + string(e.LocationColumn) + ")"
// }

package practicefundamentals

import (
	"errors"
	"fmt"
)

// ErrFixedDigit - errors.New("can't update fixed digits
// ErrhorizontalRow - The horizontal row it’s placed in already has the same digit
// ErrverticalRow - The vertical column it’s placed in has the same digit
// Errsubregion - The 3 × 3 subregion it’s placed in has the same digit
var (
	ErrFixedDigit     = errors.New("can't update a location with fixed digits")
	ErrHorizontalRow  = errors.New("horizontal row it’s placed in already has the same digit")
	ErrVerticalColumn = errors.New("vertical column it’s placed in already has the same digit")
	ErrSubregion      = errors.New("3 × 3 subregion already has the same digit")
)

//SodukoGrid is a fixed 9x9 soduko Grid array.  Firs letter capitalized means we plan to export it
type SodukoGrid struct {
	Grid [9][9]int8
	Fd   [9][9]bool //Fixed Digits will be marked as true
}

//NewSudokuGrid is constructor
func NewSudokuGrid(s [9][9]int8) *SodukoGrid {
	var sg SodukoGrid
	sg.Grid = s
	for i, r := range sg.Grid {
		for j, c := range r {
			if c != 0 {
				sg.Fd[i][j] = true
			}
		}
	}
	//fmt.Println("New Suduko Grid Constructed: ")
	//sg.prettyPrint()
	return &sg
}

// set a digit at a specific location. This method returns an error if placing the digit breaks one of the rules.
// The digit being placed may not already appear in any of the following:
// The horizontal row it’s placed in
// The vertical column it’s placed in
// The 3 × 3 subregion it’s placed in
func (sg *SodukoGrid) set(digit, locationRow, locationColumn int8) error {

	fmt.Printf("Set digit: %v at location(%v, %v)\n", digit, locationRow, locationColumn)

	//Error if the location is not zero
	if sg.Fd[locationRow][locationColumn] {
		return ErrFixedDigit
	}

	//Error if digit being placed alreayd exisits in the horizontal row it’s placed in.
	//range sg = going through every row, range sg[locationRow] = going through every column in that row.
	//Since you want to make sure this digit is nowhere else in that row, you have to go through every column in that row
	for _, v := range sg.Grid[locationRow] { //Go through every column in that row.
		if v == digit {
			return ErrHorizontalRow
		}
	}

	//Error if digit being placed alreayd exisits in the vertical column it’s placed in
	for _, v := range sg.Grid {
		if v[locationColumn] == digit {
			return ErrVerticalColumn
		}
	}

	//Error if digit being placed alreayd exisits in the 3 × 3 subregion it’s placed in
	//Define 3x3 sub region
	//For every location, the row for that sub region starts at locationRow/3
	//For every location, the column for tht sub region starts at locationColumn/3

	for row := locationRow / 3; row < (locationRow/3 + 3); row++ {
		for column := locationColumn / 3; column < locationColumn/3+3; column++ {
			if sg.Grid[row][column] == digit {
				// fmt.Printf("locationRow/3: %v\n", locationRow/3)
				// fmt.Printf("locationColumn/3: %v\n", locationColumn/3)
				return ErrSubregion
			}

		}
	}

	sg.Grid[locationRow][locationColumn] = digit
	//fmt.Println("Set(): Suduko Grid Updated: ")
	//sg.prettyPrint()

	return nil
}

//clear Method
func (sg *SodukoGrid) clear(locationRow, locationColumn int8) error {
	fmt.Printf("Clear digit at location(%v, %v)\n", locationRow, locationColumn)

	//Error if the location is not zero
	if sg.Fd[locationRow][locationColumn] {
		return ErrFixedDigit
	}
	sg.Grid[locationRow][locationColumn] = 0
	//fmt.Println("Clear(): Suduko Grid Updated: ")
	//sg.prettyPrint()

	return nil
}

func (sg SodukoGrid) prettyPrint() {
	for _, v := range sg.Grid {
		fmt.Println(v)
	}

	// for _, v := range sg.Fd {
	// 	fmt.Println(v)
	// }
}
