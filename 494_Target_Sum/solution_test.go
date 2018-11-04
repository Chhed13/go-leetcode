package _494_Target_Sum

import (
    "testing"
)

type testCase struct {
    nums  []int
    S int
    answer int
}

var testCases = []testCase{
    {
        []int{1, 1, 1, 1, 1}, 3, 5,
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := findTargetSumWays(tc.nums, tc.S)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}