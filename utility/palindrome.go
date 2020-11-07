package utility

import (
	"strings"
)

func isPalindrome(s string) bool {
	string := strings.ReplaceAll(s, " ", "")
	println(string)
	return string == reverse(string)
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes) - 1
	mid := getMidIndex(n)

	for i := 0; i < mid; i++ {
		runes[i], runes[n-i] = runes[n-i], runes[i]
	}
	return string(runes)
}

func getMidIndex(length int) int {
	mid := length / 2
	if isEven(length) {
		return mid
	}
	return mid + 1
}
