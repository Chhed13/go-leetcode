package _9_Palindrome_Number

import (
	"testing"
)

type testCase struct {
	tcase  int
	answer bool
}

var testCases = []testCase{
	{
		121, true,
	},
	{
		-121, false,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := isPalindrome(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}