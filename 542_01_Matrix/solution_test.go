package _542_01_Matrix

import (
    "testing"
)

type testCase struct {
    tcase  [][]int
    answer [][]int
}

var testCases = []testCase{
    {
        [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
        [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
    },
    {
        [][]int{{0, 0, 0}, {0, 1, 0}, {1,1,1}},
        [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := updateMatrix(tc.tcase)
        for i := range tc.answer {
            for j := range tc.answer[i] {
                if answer[i][j] != tc.answer[i][j] {
                    t.Errorf("expected: %v, got: %v", tc.answer, answer)
                }
            }
        }

    }
}
