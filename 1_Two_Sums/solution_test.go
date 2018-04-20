package _1_Two_Sums

import (
	"testing"
)

type testCase struct {
	nums   []int
	target int
	answer []int
}

var testCases = []testCase{
	{

		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := twoSum(tc.nums, tc.target)
		if len(answer) != 2 || answer[0] != tc.answer[0] || answer[1] != tc.answer[1]{
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
