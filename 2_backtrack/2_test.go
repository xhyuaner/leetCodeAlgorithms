package backtrack

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	t.Log(LetterCombinations("23")) // ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
	t.Log(LetterCombinations("2"))  // ["a", "b", "c"]
	t.Log(LetterCombinations(""))   // []
	t.Log(LetterCombinations("1"))  // []
	t.Log(LetterCombinations("37")) // ["sd", "se", "sf", "td", "te", "tf", "vd", "ve", "vf"]
	t.Log(LetterCombinations("0"))  // []
}

func TestGenerateBracket(t *testing.T) {
	t.Log(generateBracket(3)) // [((())) (()()) (())() ()(()) ()()()]
}
