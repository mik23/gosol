package utility

import (
	"bufio"
	"os"
)

/*
 * Utility to handle files
 * It provides HasNext, Next to look and read over big files
 */

type FileReader struct {
	currentString string
	reader        bufio.Reader
}

type Reader interface {
	Next() string
	HasNext() bool
}

func (fileReader *FileReader) HasNext() bool {
	bytes, _ := fileReader.reader.Peek(1)
	return len(bytes) > 0
}

func (fileReader *FileReader) Next() string {
	if !fileReader.HasNext() {
		return ""
	}
	fileReader.currentString, _ = fileReader.reader.ReadString('\n')
	return fileReader.currentString
}

func NewReader(filePath string) *FileReader {
	file, _ := os.Open(filePath)
	return newReader_(file)
}

func newReader_(file *os.File) *FileReader {
	reader := bufio.NewReader(file)
	return &FileReader{"", *reader}
}
