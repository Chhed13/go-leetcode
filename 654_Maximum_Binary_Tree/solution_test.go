package _654_Maximum_Binary_Tree

import (
	"testing"
)

type testCase struct {
	tcase  []int
	answer []int
}

var testCases = []testCase{
	{
		tcase:  []int{3, 2, 1, 6, 0, 5,},
		answer: []int{6, 3, 5, 2, 0, 1},
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		root := constructMaximumBinaryTree(tc.tcase)
		q := queue{root}
		for i, n := 0, q.Pull(); n != nil; i, n = i+1, q.Pull() {
			if n.Val != tc.answer[i] {
				t.Errorf("expected: %v, got: %v", tc.answer, n)
			}
			if n.Left != nil {
				q.Push(n.Left)
			}
			if n.Right != nil {
				q.Push(n.Right)
			}
		}
	}
}

type queue []*TreeNode

func (q *queue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *queue) Pull() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}
