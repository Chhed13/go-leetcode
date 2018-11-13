package _125_Valid_Palindrome

import (
    "testing"
)

type testCase struct {
    tcase  string
    answer bool
}

var testCases = []testCase{
    {"race a car", false},
    {"A man, a plan, a canal: Panama", true},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := isPalindrome(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
