// The 'package main' line declares that this Go file will be the starting point of the program.
package main

// Importing necessary packages for input/output and formatting.
import (
	"fmt"
	"os"
)

// The 'main' function is where the program starts its execution.
func main() {
	// Check if the number of command-line arguments is not equal to 10.
	if len(os.Args) != 10 {
		// If not equal to 10, print "Error" and exit the program.
		fmt.Println("Error")
		return
	}

	// Create a 2D grid to represent a Sudoku puzzle.
	grid := make([][]int, 9)

	// Loop to initialize each row in the grid.
	for m := range grid {
		// Initialize each row with another loop for each column.
		grid[m] = make([]int, 9)
		for n := range grid[m] {
			// Convert the character in the command-line argument to an integer and store it in the grid.
			if os.Args[m+1][n] == '.' {
				grid[m][n] = 0
			} else {
				grid[m][n] = int(os.Args[m+1][n] - '0')
			}
		}
	}

	// Call the 'solveSudoku' function to attempt solving the Sudoku puzzle.
	if solveSudoku(grid) {
		// If the puzzle is solved, print the solved grid.
		printGrid(grid)
	} else {
		// If the puzzle cannot be solved, print "Error".
		fmt.Println("Error")
	}
}

// Function to solve the Sudoku puzzle using backtracking.
func solveSudoku(grid [][]int) bool {
	// Variables to track the current row, column, and whether there are unassigned cells.
	row, col, unassigned := 0, 0, false

	// Loop to find an unassigned cell in the grid.
	for m := 0; m < 9; m++ {
		for n := 0; n < 9; n++ {
			if grid[m][n] == 0 {
				// Found an unassigned cell, store its coordinates and set 'unassigned' to true.
				row, col, unassigned = m, n, true
				break
			}
		}
		// If an unassigned cell is found, exit the outer loop.
		if unassigned {
			break
		}
	}

	// If there are no unassigned cells, the puzzle is solved.
	if !unassigned {
		return true
	}

	// Try numbers 1 through 9 in the current cell and recursively solve the remaining puzzle.
	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid) {
				return true
			}
			// If the current attempt fails, backtrack by resetting the current cell to 0.
			grid[row][col] = 0
		}
	}

	// If no number can be placed in the current cell, backtrack.
	return false
}

// Function to check if it's safe to place a number in a given cell.
func isSafe(grid [][]int, row int, col int, num int) bool {
	// Check if the number 'num' is not present in the current row, column, or 3x3 grid.
	for m := 0; m < 9; m++ {
		if grid[row][m] == num || grid[m][col] == num || grid[3*(row/3)+m/3][3*(col/3)+m%3] == num {
			return false
		}
	}
	// If 'num' is not present in the row, column, or grid, it's safe to place it.
	return true
}

// Function to print the Sudoku grid.
func printGrid(grid [][]int) {
	// Loop through each row and column, printing the numbers in the grid.
	for m := 0; m < 9; m++ {
		for n := 0; n < 9; n++ {
			fmt.Printf("%d", grid[m][n])
			// Print a space if it's not the last column.
			if n < 8 {
				fmt.Print(" ")
			}
		}
		// Move to the next line after printing each row.
		fmt.Println()
	}
}
