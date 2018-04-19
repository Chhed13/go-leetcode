package _396_Rotate_Function

import (
	"testing"
)

type testCase struct {
	tcase  []int
	answer int
}

var testCases = []testCase{
	{
		[]int{}, 0,
	},
	{
		[]int{4, 3, 2, 6}, 26,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := maxRotateFunction(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
