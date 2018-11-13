package _841_Keys_and_Rooms

import (
    "testing"
)

type testCase struct {
    tcase  [][]int
    answer bool
}

var testCases = []testCase{
    {[][]int{{1}, {2}, {3}, {}}, true,},
    {[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false,},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := canVisitAllRooms(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
