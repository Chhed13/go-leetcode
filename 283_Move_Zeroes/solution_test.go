package _283_Move_Zeroes

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer []int
}

var testCases = []testCase{
    {[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0},},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := moveZeroes(tc.tcase)
        for i := range answer {
            if answer[i] != tc.answer[i] {
                t.Errorf("expected: %v, got: %v", tc.answer, answer)
            }
        }
    }
}
