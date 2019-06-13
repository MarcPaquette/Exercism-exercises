package scrabble

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleMultiTest struct {
	input      string
	multiplier int
	expected   int
}

var scrabbleScoreMultiTests = []scrabbleMultiTest{
	{"a", 2, 2},                            // lowercase letter
	{"A", 3, 3},                            // uppercase letter
	{"f", 2, 8},                            // valuable letter
	{"at", 2, 4},                           // short word
	{"zoo", 3, 36},                         // short, valuable word
	{"street", 3, 18},                      // medium word
	{"quirky", 2, 44},                      // medium, valuable word
	{"OxyphenButazone", 3, 123},            // long, mixed-case word
	{"pinata", 2, 16},                      // english-like word
	{"", 2, 0},                             // empty input
	{"abcdefghijklmnopqrstuvwxyz", 2, 174}, // entire alphabet available
}
