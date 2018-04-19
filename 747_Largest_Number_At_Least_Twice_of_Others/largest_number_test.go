package _747_Largest_Number_At_Least_Twice_of_Others

import "testing"

type testCase struct {
	tcase  []int
	answer int
}

var testCases = []testCase{
	{
		[]int{3, 6, 1, 0,}, 1,
	},
	{
		[]int{1, 2, 3, 4,}, -1,
	},
	{
		[]int{0, 0, 3, 2,}, -1,
	},
}

func TestDominantIndex(t *testing.T) {
	for _, tc := range testCases {
		answer := dominantIndex(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
