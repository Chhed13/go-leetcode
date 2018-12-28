package _581_Shortest_Unsorted_Continuous_Subarray

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer int
}

var testCases = []testCase{
    {[]int{2, 6, 4, 8, 10, 9, 15}, 5,},
    {[]int{1, 2, 3, 4}, 0},
    {[]int{1, 3, 2, 2, 2}, 4},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := findUnsortedSubarray(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
func TestSolution2(t *testing.T) {
    for _, tc := range testCases {
        answer := findUnsortedSubarray2(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
