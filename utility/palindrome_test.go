package utility

import "testing"

func TestReverse(t *testing.T) {
	t.Run("Odd length strings should be reversed correctly", func(t *testing.T) {
		const string = "michael"
		const expected = "leahcim"
		actual := reverse(string)
		if actual != expected {
			t.Error("expected: ", expected, " actual: ", actual)
		}
	})

	t.Run("Even length strings should be reversed correctly", func(t *testing.T) {
		const string = "davide"
		const expected = "edivad"
		actual := reverse(string)
		if actual != expected {
			t.Error("expected: ", expected, " actual: ", actual)
		}
	})
}

func TestPalindrome(t *testing.T) {
	t.Run("Correct palindrome", func(t *testing.T) {
		const string = "anna"
		if !isPalindrome(string) {
			t.Error(string, " should be palindrome")
		}
	})

	t.Run("Correct palindrome sentence", func(t *testing.T) {
		const string = "eco vana voce"
		if !isPalindrome(string) {
			t.Error(string, " should be palindrome")
		}
	})

	t.Run("Incorrect palindrome", func(t *testing.T) {
		const string = "michael"
		if isPalindrome(string) {
			t.Error(string, " should not be palindrome")
		}
	})
}

func TestPalindrome2(t *testing.T) {

	t.Run("Correct palindrome ", func(t *testing.T) {
		const string = "anna"
		if !isPalindrome2(string) {
			t.Error(string, " should be palindrome")
		}
	})

	t.Run("Correct palindrome sentence", func(t *testing.T) {
		const string = "eco vana voce"
		if !isPalindrome2(string) {
			t.Error(string, " should be palindrome")
		}
	})

	t.Run("Incorrect palindrome", func(t *testing.T) {
		const string = "michael"
		if isPalindrome2(string) {
			t.Error(string, " should not be palindrome")
		}
	})
}
