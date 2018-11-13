package _98_Validate_Binary_Search_Tree

import (
    "testing"
)

func generate1() *TreeNode {
    n1 := &TreeNode{
        Val:   1,
        Left:  nil,
        Right: nil,
    }

    n2 := &TreeNode{
        Val:   3,
        Left:  nil,
        Right: nil,
    }

    n3 := &TreeNode{
        Val:   2,
        Left:  n1,
        Right: n2,
    }

    return n3
}

func generate2() *TreeNode {
    n1 := &TreeNode{
        Val:   1,
        Left:  nil,
        Right: nil,
    }

    n2 := &TreeNode{
        Val:   3,
        Left:  nil,
        Right: nil,
    }

    n3 := &TreeNode{
        Val:   6,
        Left:  n1,
        Right: n2,
    }

    n4 := &TreeNode{
        Val:   4,
        Left:  n2,
        Right: n3,
    }

    n5 := &TreeNode{
        Val:   5,
        Left:  n1,
        Right: n4,
    }

    return n5
}

func TestSolution(t *testing.T) {
    if !isValidBST(generate1()) {
        t.Errorf("case 1 expected: true")
    }
    if isValidBST(generate2()) {
        t.Errorf("case 2 expected: false")
    }
}
