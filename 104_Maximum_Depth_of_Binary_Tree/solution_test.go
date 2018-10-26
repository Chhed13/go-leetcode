package _104_Maximum_Depth_of_Binary_Tree

import (
    "testing"
)

var answer = 3

func generate() *TreeNode {
    n1 := &TreeNode{
        Val:   15,
        Left:  nil,
        Right: nil,
    }

    n2 := &TreeNode{
        Val:   7,
        Left:  nil,
        Right: nil,
    }

    n3 := &TreeNode{
        Val:   9,
        Left:  nil,
        Right: nil,
    }

    n4 := &TreeNode{
        Val:   20,
        Left:  n1,
        Right: n2,
    }

    n5 := &TreeNode{
        Val:   3,
        Left:  n3,
        Right: n4,
    }

    return n5
}

func TestSolution(t *testing.T) {
    t1 := generate()
    a := maxDepth(t1)

    if a != answer {
        t.Errorf("expected: %v, got: %v", answer, a)
    }
}
