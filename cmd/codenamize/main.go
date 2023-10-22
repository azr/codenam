package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/azr/codenam"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: codenamize [STRING]...\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(codenam.Ize(arg))
	}
	fmt.Println()
}
