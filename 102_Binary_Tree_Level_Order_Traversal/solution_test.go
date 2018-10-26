package _102_Binary_Tree_Level_Order_Traversal

import (
    "testing"
    "fmt"
)

var answer = [][]int{
    {3},
    {9, 20},
    {15, 7},
}

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
    a := levelOrder(t1)
    if len(a) != len(answer) {
        t.Errorf("expected len %v, got %v", len(answer), len(a))
    }

    for i := range a {
        for j := range a[i] {
            fmt.Println(a[i])
            if a[i][j] != answer[i][j] {
                t.Errorf("expected: %v, got: %v", answer[i][j], a[i][j])
            }
        }
    }
}


func TestSolution2(t *testing.T) {
    a := levelOrder(nil)
    if len(a) > 0 {
        t.Errorf("expected len %v, got %v", 0, len(a))
    }
}
