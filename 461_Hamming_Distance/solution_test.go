package _461_Hamming_Distance

import (
    "testing"
)

type testCase struct {
    x, y   int
    answer int
}

var testCases = []testCase{
    {4, 1, 2},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := hammingDistance(tc.x, tc.y)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
