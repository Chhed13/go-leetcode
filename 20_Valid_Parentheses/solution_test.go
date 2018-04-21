package _20_Valid_Parentheses

import (
	"testing"
)

type testCase struct {
	tcase  string
	answer bool
}

var testCases = []testCase{
	{
		"()", true,
	},
	{
		"()[]{}", true,
	},
	{
		"(())", true,
	},
	{
		"(]", false,
	},
	{
		"([)]", false,
	},
	{
		"{[]}", true,
	},
	{
		"}", false,
	},
	{
		"[", false,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := isValid(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}