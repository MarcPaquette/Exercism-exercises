package scrabble

import "testing"

func TestScoreMulti(t *testing.T) {
	for _, test := range scrabbleScoreMultiTests {
		if actual := ScoreWordMultiWord(test.input, test.multiplier); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScoreMulti(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreMultiTests {
			ScoreWordMultiWord(test.input, test.multiplier)
		}
	}
}
