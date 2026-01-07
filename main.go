package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var printLineBreaks, printWords, printChars, printBytes bool

	flag.BoolVar(&printLineBreaks, "l", false, "print line breaks")
	flag.BoolVar(&printWords, "w", false, "print words")
	flag.BoolVar(&printChars, "m", false, "print chars")
	flag.BoolVar(&printBytes, "c", false, "print bytes")

	filename := flag.CommandLine.Arg(0)
	stat, err := os.Stdin.Stat()
	

	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
}
