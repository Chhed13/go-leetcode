package _234_Palindrome_Linked_List

import (
	"testing"
)

type testCase struct {
	tcase  []int
	answer bool
}

var testCases = []testCase{
	{
		[]int{1,},
		true,
	},
	{
		[]int{1, 2, 2, 1},
		true,
	},
	{
		[]int{1, 2, 3, 2, 1},
		true,
	},
	{
		[]int{},
		true,
	},
	{
		[]int{1, 2},
		false,
	},
	{
		[]int{7, 5, 6, 7},
		false,
	},
	{
		[]int{7, 5, 6, 6, 5, 7},
		true,
	},
}

func generate(index []int) *ListNode {
	var prev *ListNode
	for i := len(index) - 1; i >= 0; i-- {
		prev = &ListNode{
			Val:  index[i],
			Next: prev,
		}
	}
	return prev
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		l1 := generate(tc.tcase)

		answer := isPalindrome(l1)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v, case: %+v", tc.answer, answer, tc.tcase)
		}
	}
}
