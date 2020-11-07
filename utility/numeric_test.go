package utility

import "testing"

func TestEven(t *testing.T) {
	t.Run("Multiples of 2 should be even", func(t *testing.T) {
		const evenNumber int = 88
		if !isEven(evenNumber) {
			t.Error("Odd number: ", evenNumber)
		}
	})
}
