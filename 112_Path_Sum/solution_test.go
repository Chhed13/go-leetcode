package _112_Path_Sum

import (
    "testing"
)

var answer1 = 22

func generate1() *TreeNode {
    n1 := &TreeNode{7,nil,nil}
    n2 := &TreeNode{2,nil,nil}
    n3 := &TreeNode{1,nil,nil}
    n4 := &TreeNode{11,n1,n2}
    n5 := &TreeNode{13,nil,nil}
    n6 := &TreeNode{4,nil,n3}
    n7 := &TreeNode{4,n4,nil}
    n8 := &TreeNode{8,n5,n6}
    n9 := &TreeNode{5,n7,n8}

    return n9
}

var answer2 = -5

func generate2() *TreeNode {
    n1 := &TreeNode{-3,nil,nil}
    n2 := &TreeNode{-2,nil,n1}

    return n2
}


func TestSolution(t *testing.T) {
    if !hasPathSum(generate1(), answer1) {
        t.Errorf("expected: true")
    }

    if !hasPathSum(generate2(), answer2) {
        t.Errorf("expected: true")
    }

}