// package pyramidpattern
package main

import (
	"fmt"
)

func main() {
	var rows int

	// Get user input
	fmt.Print("Enter number of rows: ")
	_, err := fmt.Scan(&rows)

	// Validate input
	if err != nil || rows <= 0 {
		fmt.Println("Please enter a positive integer!")
		return
	}

	// Print pyramid pattern
	fmt.Printf("\nPyramid pattern for %d rows:\n\n", rows)

	for i := 1; i <= rows; i++ {
		// Print spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// Print stars
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}

		// Move to next line
		fmt.Println()
	}
}
