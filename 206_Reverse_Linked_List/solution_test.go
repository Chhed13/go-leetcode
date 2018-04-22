package _206_Reverse_Linked_List

import (
	"testing"
	"fmt"
)

type testCase struct {
	tcase  []int
	answer []int
}

var testCases = []testCase{
	{
		[]int{1,},
		[]int{1,},
	},
	{
		[]int{1, 2, 4, 5},
		[]int{5, 4, 2, 1},
	},
	{
		[]int{},
		[]int{},
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

		answer := reverseList(l1)
		for a, i := answer, 0; i < len(tc.answer); a, i = a.Next, i+1 {
			if a == nil {
				t.Fatalf("expected equal lenght of %v", len(tc.answer))
			}
			fmt.Println(a.Val)
			if a.Val != tc.answer[i] {
				t.Errorf("expected: %v, got: %v", tc.answer[i], a.Val)
			}
		}
	}
}
