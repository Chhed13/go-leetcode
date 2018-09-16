package _145_Binary_Tree_Postorder_Traversal

/*https://leetcode.com/problems/binary-tree-postorder-traversal/description/
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func postorderTraversal1(root *TreeNode) []int {
    out := []int{}
    if root.Left != nil {
        out = append(out, postorderTraversal1(root.Left)...)
    }
    if root.Right != nil {
        out = append(out, postorderTraversal1(root.Right)...)
    }
    out = append(out, root.Val)

    return out
}

func postorderTraversal2(root *TreeNode) []int {
    s := []*TreeNode{}

    push := func(n *TreeNode) {
        s = append(s, n)
    }

    pop := func() *TreeNode {
        l := len(s)

        if l == 0 {
            return nil
        }

        n := s[l-1]
        s = s[:l-1]

        return n
    }

    out := []int{}
    rAppend := func(i int) {
        out = append([]int{i}, out...)
    }

    if root == nil {
        return out
    }

    push(root)
    for len(s) != 0 {
        n := pop()
        rAppend(n.Val)

        if n.Left != nil {
            push(n.Left)
        }
        if n.Right != nil {
            push(n.Right)
        }
    }

    return out
}
