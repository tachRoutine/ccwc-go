package wc

import "os"

type FileStats struct {
	bytes          int64
	lineBreakCount int64
	wordCount      int64
	charsCount     int64
}

func GetFileStats(file *os.File) FileStats {
	return FileStats{
		bytes: 0,
		lineBreakCount: 0,
		wordCount: 0,
		charsCount: 0,
	}
}