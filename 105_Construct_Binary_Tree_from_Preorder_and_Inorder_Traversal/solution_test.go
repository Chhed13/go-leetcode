package _105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

import (
    "testing"
)

type testCase struct {
    preorder []int
    inorder  []int
    root     int
}

var testCases = []testCase{
    {
        []int{3, 9, 20, 15, 7},
        []int{9, 3, 15, 20, 7},
        3,
    },
    {
        []int{9,},
        []int{9,},
        9,
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        if buildTree(tc.preorder, tc.inorder).Val != tc.root {
            t.Errorf("expected %d", tc.root)
        }
    }
}
