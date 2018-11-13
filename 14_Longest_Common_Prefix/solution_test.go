package _14_Longest_Common_Prefix

import (
    "testing"
)

type testCase struct {
    tcase  []string
    answer string
}

var testCases = []testCase{
    {
        []string{"flower", "flow", "flight"}, "fl",
    },
    {
        []string{"dog", "racecar", "car"}, "",
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := longestCommonPrefix(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}

func TestSolution2(t *testing.T) {
    for _, tc := range testCases {
        answer := longestCommonPrefix2(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
