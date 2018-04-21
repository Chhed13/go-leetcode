package _28_Implement_strStr

import (
	"testing"
)

type testCase struct {
	haystack string
	needle   string``
	answer   int
}

var testCases = []testCase{
	{
		"hello", "ll", 2,
	},
	{
		"aaaa", "baa", -1,
	},
	{
		"", "", 0,
	},
	{
		"a","a", 0,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := strStr(tc.haystack, tc.needle)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
