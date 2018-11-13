package _122_Best_Time_to_Buy_and_Sell_Stock_II

import (
    "testing"
)

type testCase struct {
    tcase  []int
    answer int
}

var testCases = []testCase{
    {[]int{7, 1, 5, 3, 6, 4}, 7,},
    {[]int{1, 2, 3, 4, 5}, 4,},
    {[]int{7, 6, 4, 3, 1}, 0,},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := maxProfit(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
