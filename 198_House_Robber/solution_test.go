package _198_House_Robber

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer int
}

var testCases = []testCase{
    {[]int{2, 1, 1, 2}, 4},
    {[]int{1, 2, 3, 1}, 4},
    {[]int{2, 7, 9, 3, 1}, 12},
    {[]int{1, 1, 1, 2, 2, 2}, 5},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := rob(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
