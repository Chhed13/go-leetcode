package _26_Remove_Duplicates_from_Sorted_Array

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
		[]int{1,}, 1,
	},
	{
		[]int{1,1,2}, 2,
	},
	{
		[]int{0,0,1,1,1,2,2,3,3,4,4,4}, 5,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := removeDuplicates(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}