package main

import (
	"fmt"
	"os"

	"github.com/ekilie/vintpm.git/toolkit"
)

func main() {
	args := os.Args
	if len(args)<2{
		printDefault()
	}

	switch args[1] {
	case "update-vint":
		fmt.Println("Updating vint")
		toolkit.Update()
		// break
	default:
		printDefault()
	}
	
	
}

func printDefault(){
	fmt.Println("The official vint package manager (vintpm)")

}
