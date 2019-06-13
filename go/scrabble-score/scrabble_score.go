// Package scrabble provides the scrable score for a given word
package scrabble

import (
	"strings"
)

var wordList = map[rune]int{
	'A': 1,
	'B': 3,
	'C': 3,
	'D': 2,
	'E': 1,
	'F': 4,
	'G': 2,
	'H': 4,
	'I': 1,
	'J': 8,
	'K': 5,
	'L': 1,
	'M': 3,
	'N': 1,
	'O': 1,
	'P': 3,
	'Q': 10,
	'R': 1,
	'S': 1,
	'T': 1,
	'U': 1,
	'V': 4,
	'W': 4,
	'X': 8,
	'Y': 4,
	'Z': 10,
}

// Score gives the scrabble score for a given word
func Score(word string) (score int) {

	return ScoreWordMultiLetter(word, nil)

}

// ScoreWordMultiWord gives the scrabble score for a double or triple word for a given word
func ScoreWordMultiWord(word string, multiplier int) (score int) {

	// this should create an array with the multiplier at each position, but this is easier and more readable
	score = ScoreWordMultiLetter(word, nil)
	return score * multiplier
}

// ScoreWordMultiLetter gives the scrabble score for a given word with a triple or double letter in an positional array.
//  The multiplerPositions array should have the index location of the mulitipler and the multiper value of 2 or 3 placed inside.
//  For example:  word = "entail" and multiplierPositions = [1, 1, 1, 3, 1, 2] would have the "A" tripled and the "L" doubled.
func ScoreWordMultiLetter(word string, multiplierPositions []int) (score int) {
	if word == "" {
		return 0
	}

	for i, letter := range strings.ToUpper(word) {
		if i < len(multiplierPositions) {
			if (multiplierPositions[i] > 3) || (multiplierPositions[i] < 0) {
				return 0
			}
			score += wordList[letter] * multiplierPositions[i]
		} else {
			score += wordList[letter]
		}
	}

	return score
}
