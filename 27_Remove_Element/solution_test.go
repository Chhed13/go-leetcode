package _27_Remove_Element

import (
	"testing"
)

type testCase struct {
	val    []int
	num    int
	answer int
}

var testCases = []testCase{
	{
		[]int{3,2,2,3}, 3,2,
	},
	{
		[]int{0,1,2,2,3,0,4,2}, 2, 5,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := removeElement(tc.val,tc.num)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
