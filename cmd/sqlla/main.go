package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
	"github.com/mackee/go-sqlla/v2"
)

var Version string

func main() {
	var opts sqlla.Options
	kong.Parse(&opts)

	if opts.Version {
		fmt.Println("sqlla - Type safe, reflect free, generative SQL Builder + ORM-like methods")
		fmt.Printf("version %s\n", Version)
		os.Exit(0)
	}

	if err := sqlla.Run(opts); err != nil {
		slog.Error("occurred error", slog.Any("error", err))
		os.Exit(1)
	}
}
