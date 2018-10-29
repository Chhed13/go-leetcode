package _279_Perfect_Squares

import (
    "testing"
)

type testCase struct {
    tcase  int
    answer int
}

var testCases = []testCase{
    {12, 3,},
    {13, 2},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := numSquares(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
