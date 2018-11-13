package _217_Contains_Duplicate

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer bool
}

var testCases = []testCase{
    {[]int{1, 2, 3, 1,}, true,},
    {[]int{1, 2, 3, 4,}, false,},
    {[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2,}, true,},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := containsDuplicate(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
