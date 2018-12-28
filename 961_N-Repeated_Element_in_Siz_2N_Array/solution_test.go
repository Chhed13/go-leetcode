package _961_N_Repeated_Element_in_Siz_2N_Array

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer int
}

var testCases = []testCase{
    {[]int{1, 2, 3, 3}, 3,},
    {[]int{2, 1, 2, 5, 3, 2}, 2},
    {[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := repeatedNTimes(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}

func TestSolution2(t *testing.T) {
    for _, tc := range testCases {
        answer := repeatedNTimes2(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
