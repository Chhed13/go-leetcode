package _553_Optimal_Division

import (
	"testing"
)

type testCase struct {
	tcase  []int
	answer string
}

var testCases = []testCase{
	{
		[]int{10, 2, 3},
		"10/(2/3)",
	},
	{
		[]int{1000, 100, 10, 2},
		"1000/(100/10/2)",
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := optimalDivision(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
