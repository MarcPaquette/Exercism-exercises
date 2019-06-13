package scrabble

import "testing"

func TestScoreMultiLetter(t *testing.T) {
	for _, test := range scrabbleScoreMultiLetterTests {
		if actual := ScoreWordMultiLetter(test.input, test.multiplierArray); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkScoreMultiLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreMultiLetterTests {
			ScoreWordMultiLetter(test.input, test.multiplierArray)
		}
	}
}
