package _387_First_Unique_Character_in_a_String

import (
    "testing"
)

type testCase struct {
    tcase  string
    answer int
}

var testCases = []testCase{
    {"leetcode", 0,},
    {"loveleetcode", 2,},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := firstUniqChar(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
