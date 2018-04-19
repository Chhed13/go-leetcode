package _675_Cut_Off_Trees_for_Golf_Event

import "testing"

type tcase struct {
	tcase  [][]int
	answer int
}

var testCases = []tcase{
	{
		[][]int{
			{1, 2, 3},
			{0, 0, 4},
			{7, 6, 5},
		}, 6,
	},
	{
		[][]int{
			{1, 2, 3},
			{0, 0, 0},
			{7, 6, 5},
		}, -1,
	},
	{
		[][]int{
			{2, 3, 4},
			{0, 0, 5},
			{8, 7, 6},
		}, 6,
	},
	{
		[][]int{
			{0, 0, 0, 3528, 2256, 9394, 3153},
			{8740, 1758, 6319, 3400, 4502, 7475, 6812},
			{0, 0, 3079, 6312, 0, 0, 0},
			{6828, 0, 0, 0, 0, 0, 8145},
			{6964, 4631, 0, 0, 0, 4811, 0},
			{0, 0, 0, 0, 9734, 4696, 4246},
			{3413, 8887, 0, 4766, 0, 0, 0},
			{7739, 0, 0, 2920, 0, 5321, 2250},
			{3032, 0, 3015, 0, 3269, 8582, 0},
		}, -1,
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		answer := cutOffTree(tc.tcase)
		if answer != tc.answer {
			t.Errorf("wrong answer. Expected: %d, got: %d, case: %d", tc.answer, answer, i)
		}
	}
}
