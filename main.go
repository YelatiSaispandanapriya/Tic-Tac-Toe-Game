package main

import (
	"fmt"
)

// Constants for players
const (
	Empty = " "
	X     = "X"
	O     = "O"
)

// Board represents the Tic Tac Toe grid
type Board [3][3]string

// Initialize the board with empty spaces
func NewBoard() Board {
	return Board{
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
	}
}

// Display the board
func (b Board) Display() {
	for i, row := range b {
		fmt.Printf(" %s | %s | %s \n", row[0], row[1], row[2])
		if i < 2 {
			fmt.Println("-----------")
		}
	}
}

// Check if a player has won
func (b Board) CheckWin(player string) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if b[i][0] == player && b[i][1] == player && b[i][2] == player {
			return true // Row win
		}
		if b[0][i] == player && b[1][i] == player && b[2][i] == player {
			return true // Column win
		}
	}

	// Check diagonals
	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return true // Diagonal win
	}
	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return true // Diagonal win
	}

	return false
}

// Check if the board is full (draw)
func (b Board) IsFull() bool {
	for _, row := range b {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

// Make a move on the board
func (b *Board) MakeMove(row, col int, player string) bool {
	// Convert 1-based indexing to 0-based
	row--
	col--

	if row < 0 || row >= 3 || col < 0 || col >= 3 || b[row][col] != Empty {
		return false // Invalid move
	}
	b[row][col] = player
	return true
}

func main() {
	board := NewBoard()
	currentPlayer := X

	for {
		// Display the board
		fmt.Println("\nCurrent Board:")
		board.Display()

		// Get player input
		var row, col int
		fmt.Printf("Player %s, enter row (1-3) and column (1-3): ", currentPlayer)
		fmt.Scan(&row, &col)

		// Make the move
		if !board.MakeMove(row, col, currentPlayer) {
			fmt.Println("Invalid move! Try again.")
			continue
		}

		// Check for a win
		if board.CheckWin(currentPlayer) {
			fmt.Println("\nFinal Board:")
			board.Display()
			fmt.Printf("Player %s wins!\n", currentPlayer)
			break
		}

		// Check for a draw
		if board.IsFull() {
			fmt.Println("\nFinal Board:")
			board.Display()
			fmt.Println("It's a draw!")
			break
		}

		// Switch players
		if currentPlayer == X {
			currentPlayer = O
		} else {
			currentPlayer = X
		}
	}
}