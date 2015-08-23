package main

import (
	"os"

	"github.com/mackee/go-sqlla"
)

func main() {
	from := os.Getenv("GOFILE")
	if from == "" {
		args := os.Args
		if len(args) == 0 {
			os.Exit(1)
		}
		from = args[0]
	}
	sqlla.Run(from)
}
