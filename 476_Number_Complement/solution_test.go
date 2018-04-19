package _476_Number_Complement

import (
	"testing"
)

type testCase struct {
	tcase  int
	answer int
}

var testCases = []testCase{
	{
		5, 2,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := findComplement(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
