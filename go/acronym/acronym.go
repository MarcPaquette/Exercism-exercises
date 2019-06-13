/*
Package acronym generates acronyms out of provided strings
*/
package acronym

import (
	"strings"
	"unicode"
)

//firstLetter gets the first letter of a string
func firstLetter(s string) string {
	if len(s) <= 0 {
		return ""
	}
	wordLetters := []rune(s)
	letter := wordLetters[0]

	if unicode.IsLetter(letter) {
		return string(letter)
	}
	return ""
}

//phraseBreaker breaks the phrase into seperate words.
func phraseBreaker(s string) []string {
	spacedPhrase := strings.ReplaceAll(s, "-", " ")
	return strings.Split(spacedPhrase, " ")

}

// Abbreviate takes a string and makes an acronym out of it.
func Abbreviate(s string) string {
	upperPhrase := strings.ToUpper(s)

	acronym := ""

	for _, word := range phraseBreaker(upperPhrase) {
		acronym += firstLetter(word)

	}

	return acronym
}
