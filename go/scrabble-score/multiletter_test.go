package scrabble

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleMultiLetterTest struct {
	input           string
	multiplierArray []int
	expected        int
}

var scrabbleScoreMultiLetterTests = []scrabbleMultiLetterTest{
	{"a", []int{1}, 1},                // lowercase letter
	{"A", []int{1}, 1},                // uppercase letter
	{"f", []int{2}, 8},                // valuable letter
	{"at", []int{1, 2}, 3},            // short word
	{"zoo", []int{3, 1, 2}, 33},       // short, valuable word
	{"street", []int{3, 2, -1, 0}, 0}, // test to see if short array  returns 0

	{"street", []int{3, 2, 1, 0, -1, 0}, 0}, // test to see if out of bounds values returns 0
}
