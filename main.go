package main

import (
	"fmt"
	"os"

	"github.com/ekilie/vintpm.git/toolkit"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "update-vint":
		fmt.Println("Updating vint")
		toolkit.Update()
		// break
	}

}
