package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mackee/go-sqlla"
)

var Version string

func main() {
	var isShowVersion bool
	flag.BoolVar(&isShowVersion, "version", false, "show this version")
	flag.Parse()

	if isShowVersion {
		fmt.Println("sqlla - Type safe, reflect free, generative SQL Builder + ORM-like methods")
		fmt.Printf("version %s\n", Version)
		os.Exit(0)
	}

	from := os.Getenv("GOFILE")
	if from == "" {
		args := os.Args
		if len(args) == 0 {
			os.Exit(1)
		}
		from = args[0]
	}
	ext := os.Getenv("SQLLA_GENERATE_FILE_EXT")
	if ext == "" {
		ext = ".gen.go"
	}
	sqlla.Run(from, ext)
}
