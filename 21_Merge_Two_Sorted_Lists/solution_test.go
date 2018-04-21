package _21_Merge_Two_Sorted_Lists

import (
	"testing"
	"fmt"
)

type testCase struct {
	tcase1 []int
	tcase2 []int
	answer []int
}

var testCases = []testCase{
	{
		[]int{1, 2, 4},
		[]int{1, 3, 4},
		[]int{1, 1, 2, 3, 4, 4},
	},
}

func generate(index []int) *ListNode {
	var prev *ListNode
	for i := len(index) - 1; i >= 0; i-- {
		prev = &ListNode{
			Val:  index[i],
			Next: prev,
		}
		fmt.Println("gen", index, prev.Val)
	}
	return prev
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		l1 := generate(tc.tcase1)
		l2 := generate(tc.tcase2)

		answer := mergeTwoLists(l1, l2)
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

func TestSolution2(t *testing.T) {
	for _, tc := range testCases {
		l1 := generate(tc.tcase1)
		l2 := generate(tc.tcase2)

		answer := mergeTwoLists2(l1, l2)
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