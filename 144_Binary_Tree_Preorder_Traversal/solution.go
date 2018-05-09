package _144_Binary_Tree_Preorder_Traversal

/*https://leetcode.com/problems/binary-tree-preorder-traversal/description/

Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal1(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	out = append(out, root.Val)
	if root.Left != nil {
		out = append(out, preorderTraversal1(root.Left)...)
	}
	if root.Right != nil {
		out = append(out, preorderTraversal1(root.Right)...)
	}
	return out
}

func preorderTraversal2(root *TreeNode) []int {
	s := []*TreeNode{}

	push := func(t *TreeNode) {
		s = append(s, t)
	}

	pop := func() *TreeNode {
		l := len(s)

		if (l == 0) {
			return nil
		}
		r := (s[l-1])

		s = s[:l-1]
		return r
	}

	out := []int{}
	if root == nil {
		return out
	}

	for n := root; n != nil; n = pop() {
		out = append(out, n.Val)
		if n.Right != nil {
			push(n.Right)
		}
		if n.Left != nil {
			push(n.Left)
		}
	}
	return out
}