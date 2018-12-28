package _771_Jewels_and_Stones

import (
    "testing"
)

type testCase struct {
    J, S   string
    answer int
}

var testCases = []testCase{
    {"aA", "aAAbbbb", 3,},
    {"z", "ZZ", 0,},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := numJewelsInStones(tc.J, tc.S)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
