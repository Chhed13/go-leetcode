package _273_Integer_to_English_Words

import "testing"

type testCase struct {
	tcase  int
	answer string
}

var testCases = []testCase{
	{
		1,
		"One",
	},
	{
		123,
		"One Hundred Twenty Three",
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := numberToWords(tc.tcase)
		if len(answer) != len(tc.answer) {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
