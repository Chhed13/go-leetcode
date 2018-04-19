package _17_Letter_Combinations_Of_Phone_Number

import (
	"testing"
)

type testCase struct {
	tcase  string
	answer []string
}

var testCases = []testCase{
	{
		"23",
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := letterCombinations(tc.tcase)
		if len(answer) != len(tc.answer) {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
