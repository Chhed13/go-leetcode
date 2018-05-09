package _94_Binary_Tree_Inorder_Traversal

import (
	"testing"
	"fmt"
)

var answer = []int{
	1, 3, 2,
}

func generate1() *TreeNode {
	n1 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	n2 := &TreeNode{
		Val:   2,
		Left:  n1,
		Right: nil,
	}

	n3 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: n2,
	}
	return n3
}

func TestSolution1(t *testing.T) {
	t1 := generate1()
	a := inorderTraversal1(t1)
	if len(a) != len(answer) {
		t.Errorf("expected len %v, got %v", len(answer), len(a))
	}

	for i := range a {
		fmt.Println(a[i])
		if a[i] != answer[i] {
			t.Errorf("expected: %v, got: %v", answer[i], a[i])
		}
	}
}

func TestSolution2(t *testing.T) {
	t1 := generate1()
	a := inorderTraversal2(t1)
	if len(a) != len(answer) {
		t.Errorf("expected len %v, got %v", len(answer), len(a))
	}

	for i := range a {
		fmt.Println(a[i])
		if a[i] != answer[i] {
			t.Errorf("expected: %v, got: %v", answer[i], a[i])
		}
	}
}