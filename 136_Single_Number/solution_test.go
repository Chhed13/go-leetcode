package _136_Single_Number

import (
	"testing"
)

type testCase struct {
	tcase  []int
	answer int
}

var testCases = []testCase{
	{
		[]int{2,2,1,}, 1,
	},
	{
		[]int{4,1,2,1,2}, 4,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := singleNumber(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}