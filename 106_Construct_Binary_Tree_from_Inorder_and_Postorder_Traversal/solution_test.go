package _106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

import (
    "testing"
)

type testCase struct {
    inorder   []int
    postorder []int
    root      int
}

var testCases = []testCase{
    {
        []int{9, 3, 15, 20, 7},
        []int{9, 15, 7, 20, 3},
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
        if buildTree(tc.inorder, tc.postorder).Val != tc.root {
            t.Errorf("expected %d", tc.root)
        }
    }
}
