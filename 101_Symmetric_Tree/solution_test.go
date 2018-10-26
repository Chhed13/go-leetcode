package _101_Symmetric_Tree

import (
    "testing"
)

func generate1() *TreeNode {
    n1 := &TreeNode{3,nil,nil}
    n2 := &TreeNode{4,nil,nil}
    n3 := &TreeNode{4,nil,nil}
    n4 := &TreeNode{3,nil,nil}
    n5 := &TreeNode{2,n2,n1}
    n6 := &TreeNode{2,n4,n3}
    n7 := &TreeNode{1,n5,n6}

    return n7
}

func generate2() *TreeNode {
    n1 := &TreeNode{3,nil,nil}
    n4 := &TreeNode{3,nil,nil}
    n5 := &TreeNode{2,n1,nil}
    n6 := &TreeNode{2,n4,nil}
    n7 := &TreeNode{1,n5,n6}

    return n7
}

func TestSolution(t *testing.T) {
    if !isSymmetric(generate1()) {
        t.Errorf("expected: true")
    }
    if isSymmetric(generate2()) {
        t.Errorf("expected: false")
    }
}

func TestSolution2(t *testing.T) {
    if !isSymmetric2(generate1()) {
        t.Errorf("expected: true")
    }
    if isSymmetric2(generate2()) {
``        t.Errorf("expected: false")
    }
}