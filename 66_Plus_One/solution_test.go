package _66_Plus_One

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer []int
}

var testCases = []testCase{
    {[]int{1, 2, 3,}, []int{1, 2, 4,}},
    {[]int{4, 3, 2, 1,}, []int{4, 3, 2, 2,}},
    {[]int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := plusOne(tc.tcase)
        for i := range answer {
            if answer[i] != tc.answer[i] {
                t.Errorf("expected: %v, got: %v", tc.answer, answer)
            }
        }
    }
}
