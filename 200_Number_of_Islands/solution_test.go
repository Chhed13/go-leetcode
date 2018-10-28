package _200_Number_of_Islands

import (
    "fmt"
    "testing"
)

type testCase struct {
    tcase  [][]byte
    answer int
}

var testCases = []testCase{
    {
        [][]byte{
            {1, 1, 1, 1, 0,},
            {1, 1, 0, 1, 0,},
            {1, 1, 0, 0, 0,},
            {0, 0, 0, 0, 0,},
        },
        1,
    }, {
        [][]byte{
            {1, 1, 0, 0, 0,},
            {1, 1, 0, 0, 0,},
            {0, 0, 1, 0, 0,},
            {0, 0, 0, 1, 1,},
        },
        3,
    }, {
        [][]byte{
            {1, 1, 0, 0, 0},
            {1, 1, 0, 0, 0},
            {0, 0, 1, 0, 0},
            {0, 0, 0, 1, 1},
        },
        3,
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := numIslands(tc.tcase)
        fmt.Println(answer)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
