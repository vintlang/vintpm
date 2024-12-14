package main

import (
	"fmt"
	"os"

	"github.com/ekilie/vintpm/toolkit"
)

// Entry point of the application
func main() {
	args := os.Args

	// Ensure there's at least one argument
	if len(args) < 2 {
		printDefault()
		return
	}

	// Handle command-line arguments
	switch args[1] {
	case "update-vint":
		fmt.Println("Updating Vint...")
		toolkit.Update()
	default:
		printDefault()
	}
}

// Displays the default help message
func printDefault() {
	fmt.Println("The official Vint package manager (vintpm)")
	fmt.Println("Usage:")
	fmt.Println("  vintpm update-vint   Update the Vint programming language.")
}
