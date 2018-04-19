package _547_Find_Circles

import "testing"

type testCase struct {
	tcase  [][]int
	answer int
}

var testCases = []testCase{
	{
		[][]int{
			{1,},
		}, 1,
	},
	{
		[][]int{
			{1, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		}, 2,
	},
	{
		[][]int{
			{1,1,0},
			{1,1,1},
			{0,1,1},
		}, 1,

	},
	{
		[][]int{
			{1,0,1,0,1},
			{0,1,0,0,0},
			{1,0,1,0,0},
			{0,0,0,1,0},
			{1,0,0,0,1},
		}, 3,

	},
	//[[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]]
	{
		[][]int {
			{1,0,0,1},
			{0,1,1,0},
			{0,1,1,1},
			{1,0,1,1},
		}, 1,
	},
}

func TestSolution(t *testing.T) {
	for _,tc := range testCases {
		answer := findCircleNum(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}
