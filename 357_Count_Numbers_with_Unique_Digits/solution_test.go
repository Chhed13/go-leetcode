package _357_Count_Numbers_with_Unique_Digits

import (
	"testing"
)

type testCase struct {
	tcase  int
	answer int
}

var testCases = []testCase{
	{
		0, 1,
	},
	{
		1, 10,
	},
	{
		2, 91,
	},
	{
		3, 739,
	},
}

func TestDominantIndex(t *testing.T) {
	for _, tc := range testCases {
		answer := countNumbersWithUniqueDigits(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
