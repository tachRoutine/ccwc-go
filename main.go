package main

import (
	"flag"
)

func main() {
	var printLineBreaks, printWords, printChars, printBytes bool

	flag.BoolVar(&printLineBreaks, "l", false, "print line breaks")
	flag.BoolVar(&printWords, "w", false, "print words")
	flag.BoolVar(&printChars, "m", false, "print chars")
	flag.BoolVar(&printBytes, "c", false, "print bytes")
	
	filename := flag.CommandLine.Arg(0)
	flag.Parse()
}
