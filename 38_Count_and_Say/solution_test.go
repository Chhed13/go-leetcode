package _38_Count_and_Say

import (
	"testing"
)

type testCase struct {
	tcase  int
	answer string
}

var testCases = []testCase{
	{
		1, "1",
	},
	{
		2, "11",
	},
	{
		3, "21",
	},
	{
		4, "1211",
	},
	{
		5, "111221",
	},
	{
		6, "312211",
	},
	{
		7, "13112221",
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := countAndSay(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}