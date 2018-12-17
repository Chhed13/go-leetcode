package _118_Pascals_Triangle

import (
    "testing"
)

type testCase struct {
    tcase  int
    answer [][]int
}

var testCases = []testCase{
    {
        5,
        [][]int{
            {1},
            {1, 1},
            {1, 2, 1},
            {1, 3, 3, 1},
            {1, 4, 6, 4, 1},
        },
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := generate(tc.tcase)
        for i := range answer {
            for j := range answer[i] {
                if answer[i][j] != tc.answer[i][j] {
                    t.Errorf("expected: %v, got: %v", tc.answer, answer)
                }
            }
        }
    }
}
