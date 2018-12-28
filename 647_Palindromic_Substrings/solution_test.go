package _647_Palindromic_Substrings

import (
    "testing"
)

type testCase struct {
    tcase  string
    answer int
}

var testCases = []testCase{
    {"abc", 3},
    {"aaa", 6},
    {"a", 1},
}

func TestSolutionSlow(t *testing.T) {
    for _, tc := range testCases {
        answer := countSubstringsSlow(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}

func TestSolutionFaster(t *testing.T) {
    for _, tc := range testCases {
        answer := countSubstringsFaster(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
