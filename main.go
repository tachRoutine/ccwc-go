package main

import (
	"ccwc/wc"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	var printLineBreaks, printWords, printChars, printBytes bool

	flag.BoolVar(&printLineBreaks, "l", false, "print line breaks")
	flag.BoolVar(&printWords, "w", false, "print words")
	flag.BoolVar(&printChars, "m", false, "print chars")
	flag.BoolVar(&printBytes, "c", false, "print bytes")

	flag.Parse()
	
	filename := flag.CommandLine.Arg(0)
	if !printLineBreaks && !printWords && !printChars && !printBytes {
		printLineBreaks = true
		printWords = true
		printBytes = true
	}

	stat, err := os.Stdin.Stat()

	if err != nil {
		log.Fatal(err)
	}

	var file *os.File
	if (stat.Mode() && os.ModeCharDevice) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	fStats := wc.GetFileStats(file)
	var cols []string

	if printLineBreaks {
		cols = append(cols, strconv.FormatInt(fStats.LineBreakCount, 10))
	}

	if printWords {
		cols = append(cols, strconv.FormatInt(fStats.WordCount, 10))
	}

	if printChars {
		cols = append(cols, strconv.FormatInt(fStats.CharsCount, 10))
	}

	if printBytes {
		cols = append(cols, strconv.FormatInt(fStats.Bytes, 10))
	}
}
