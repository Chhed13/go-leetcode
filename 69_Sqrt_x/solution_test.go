package _69_Sqrt_x

import (
	"testing"
)

type testCase struct {
	tcase  int
	answer int
}

var testCases = []testCase{
	{
		1, 1,
	},
	{
		3, 1,
	},
	{
		4, 2,
	},
	{
		8, 2,
	},
	{
		9, 3,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := mySqrt(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}