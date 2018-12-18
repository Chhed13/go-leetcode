package _268_Missing_Number

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer int
}

var testCases = []testCase{
    {[]int{3, 0, 1}, 2,},
    {[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := missingNumber(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
