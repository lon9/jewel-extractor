package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		slot int
		lang string
	)

	flag.IntVar(&slot, "s", 1, "save slot")
	flag.StringVar(&lang, "l", "en", "language of output")
	flag.Parse()

	fname := flag.Arg(0)
	if fname == "" {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	extractor, err := NewExtractor()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := extractor.Extract(fname, slot, lang); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
