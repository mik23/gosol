package utility

import (
	"testing"
)

func TestNewReader(t *testing.T) {
	t.Run("should find a byte for populated text file", func(t *testing.T) {
		fileReader:= NewReader("../resource/testfile.txt")
		if !fileReader.HasNext() {
			t.Error("Err")
		}
	})

	t.Run("should read correctly the last word in a populated text file", func(t *testing.T) {
		fileReader:= NewReader("../resource/testfile.txt")
		for fileReader.HasNext() {
			fileReader.Next()
		}

		if fileReader.currentString != "tutti!" {
			t.Error("maleee!!!!")
		}
	})
}