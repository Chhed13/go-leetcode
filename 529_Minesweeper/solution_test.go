package _529_Minesweeper

import (
	"testing"
)

type testCase struct {
	board  [][]byte
	click  []int
	answer [][]byte
}

var testCases = []testCase{
	//[[],[],[],[]]
	//[3,0]
	{
		board: [][]byte{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
		},
		click: []int{3, 0},
		answer: [][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := updateBoard(tc.board, tc.click)
		for i := 0; i < len(answer); i++ {
			for j := 0; j < len(answer[0]); j++ {
				if answer[i][j] != tc.answer[i][j] {
					t.Errorf("expected: %v, got: %v", tc.answer, answer)
				}
			}
		}
	}
}
