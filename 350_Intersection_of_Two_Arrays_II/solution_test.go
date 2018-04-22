package _350_Intersection_of_Two_Arrays_II

import (
	"testing"
)

type testCase struct {
	nums1  []int
	nums2  []int
	answer []int
}

var testCases = []testCase{
	{
		[]int{1, 2, 2, 1,},
		[]int{2, 2,},
		[]int{2, 2,},
	},
	{
		[]int{},
		[]int{},
		[]int{},
	},
	{
		[]int{1,2,1},
		[]int{},
		[]int{},
	},
	{
		[]int{1,2,1,},
		[]int{2},
		[]int{2},
	},
	{
		[]int{1,2,},
		[]int{1,1,},
		[]int{1},
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := intersect(tc.nums1, tc.nums2)
		for i := 0; i < len(answer); i++ {
			if answer[i] != tc.answer[i] {
				t.Errorf("expected: %+v, got: %+v", tc.answer, answer)
			}
		}
	}
}
