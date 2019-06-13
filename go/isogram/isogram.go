// Package isogram is a single function that identifies word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.
package isogram

import (
	"strings"
)

// IsIsogram takes a string values and returns a true if it identifies word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times
func IsIsogram(word string) bool {

	letterPresent := make(map[rune]bool)

	for _, letter := range strings.ToUpper(word) {
		//ignore hyphens & spaces
		if letter == '-' || letter == ' ' {
			continue
		}

		if letterPresent[letter] == true {
			return false
		}

		letterPresent[letter] = true
	}
	return true
}
