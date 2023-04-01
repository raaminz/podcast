package main

import (
	"flag"
	"io"
	"os"

	"github.com/raaminz/compilepodcast/cmd"
)

func main() {
	dryRun := flag.Bool("d", false, "print output without writing to file")
	flag.Parse()
	var writer io.Writer
	if *dryRun {
		writer = io.Writer(os.Stdout)
	} else {
		// create and open file
		var err error
		writer, err = os.Create("../feed.rss")
		if err != nil {
			panic(err)
		}
	}
	// generate feed
	err := cmd.GenerateRSS(writer)
	if err != nil {
		panic(err)
	}
}
