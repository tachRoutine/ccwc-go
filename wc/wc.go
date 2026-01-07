package wc

import "os"

type FileStats struct {
	Bytes          int64
	LineBreakCount int64
	WordCount      int64
	CharsCount     int64
}

func GetFileStats(file *os.File) FileStats {
	return FileStats{
		Bytes: 0,
		LineBreakCount: 0,
		WordCount: 0,
		CharsCount: 0,
	}
}