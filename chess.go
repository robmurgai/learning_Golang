package main

import "fmt"

/**
 **
 ** Experiment: chess.go
 ** Display all the chess pieces at their starting positions using the characters kqrbnp for white pieces along the bottom and
 ** uppercase KQRBNP for black pieces on the top.
 ** Write a function that nicely displays the board.
 ** Instead of strings, use [8][8]rune to represent the board. Recall that rune literals are surrounded with single quotes and can be printed with the
 ** %c format verb.
 **
 ** R N B Q K R N B
 ** P P P P P P P P
 **
 **
 **
 **
 ** p p p p p p p p
 ** r n b q k r n b
**/

func chess() {

	var board [8][8]rune

	//Black Pieces
	board[0] = [8]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}

	for i := range board[1] {
		board[1][i] = 'P'
	}

	for i := range board[6] {
		board[6][i] = 'p'
	}

	//White Pieces
	board[7] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}

	printChessBoard(board)

}

func printChessBoard(board [8][8]rune) {

	for row := range board {
		for column := range board[row] {
			if board[row][column] != 0 {
				fmt.Printf("%c ", board[row][column])
			} else {
				fmt.Printf("- ")
			}
		}
		fmt.Println()
	}

}

func emptyRuneis0() {
	var emptyRune rune
	fmt.Printf("Is emptyRune the same as 0: %v", emptyRune == 0)
}
