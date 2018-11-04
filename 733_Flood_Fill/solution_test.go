package _733_Flood_Fill

import (
    "testing"
)

type testCase struct {
    image            [][]int
    sr, sc, newColor int
    answer           [][]int
}

var testCases = []testCase{
    {
        [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1},},
        1, 1, 2,
        [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1},},
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := floodFill(tc.image,tc.sr,tc.sc,tc.newColor)
        for i := range tc.answer {
            for j := range tc.answer[i] {
                if answer[i][j] != tc.answer[i][j] {
                    t.Errorf("expected: %v, got: %v", tc.answer, answer)
                }
            }
        }

    }
}
