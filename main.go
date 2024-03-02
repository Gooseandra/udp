package main

import (
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 3 {
		client(args[0], args[1], args[2])
	} else if len(args) == 1 {
		server(args[0])
	}

}
