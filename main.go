package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]

	switch arg {
	case "update-vint":
		fmt.Println("Updating vint")
		break
	}

}
