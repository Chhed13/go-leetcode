package _7_Reverse_Integer

import (
	"testing"
)

type testCase struct {
	tcase  int
	answer int
}

var testCases = []testCase{
	{
		123, 321,
	},
	{
		-456, -654,
	},
	{
		120, 21,
	},
	{
		2147483647, 0,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := reverse(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}