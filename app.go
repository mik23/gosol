package main

import (
	"example.com/gosol/utility"
	"fmt"
)

func readFile(reader utility.Reader){
	for reader.HasNext() {
		fmt.Println(reader.Next())
	}
}

func main() {
	fileReader:= utility.NewReader("resource/testfile.txt")
	readFile(fileReader)
}