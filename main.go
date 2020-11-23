package main

import (
	"flag"
	"fmt"
)

func main() {
	handleArgs()
}

//handle input arguments
func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "db:reset":
			Reset()
		case "db:create":
			Create()
		case "db:delete":
			Delete()
		case "db:migrate":
			Migrate()
		case "db:seed":
			if len(args[1]) == 0 {
				fmt.Println("seeding all")
			} else {
				fmt.Printf("seeding %v", args[1:])
			}
		default:
			fmt.Println("argument not available")
		}
	} else {
		Start()
	}
}
