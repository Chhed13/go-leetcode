package _376_Wiggle_Subsequence

import (
	"testing"
)

type testCase struct {
	tcase  []int
	answer int
}

var testCases = []testCase{
	{
		[]int{
			1, 7, 4, 9, 2, 5,
		}, 6,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := wiggleMaxLength(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
