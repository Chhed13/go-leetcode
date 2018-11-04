package _739_Daily_Temperatures

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer []int
}

var testCases = []testCase{
    {
        []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0,},
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := dailyTemperatures(tc.tcase)
        for i := range answer {
            if answer[i] != tc.answer[i] {
                t.Errorf("expected: %v, got: %v", tc.answer[i], answer[i])
            }
        }
    }
}