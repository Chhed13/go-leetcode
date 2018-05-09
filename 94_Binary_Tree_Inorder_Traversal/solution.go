package _94_Binary_Tree_Inorder_Traversal

/*https://leetcode.com/problems/binary-tree-inorder-traversal/description/

Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	if root.Left != nil {
		out = append(out, inorderTraversal1(root.Left)...)
	}
	out = append(out, root.Val)
	if root.Right != nil {
		out = append(out, inorderTraversal1(root.Right)...)
	}
	return out
}

func inorderTraversal2(root *TreeNode) []int {
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

	n := root


	for n != nil {
		for n.Left != nil {
			push(n)
			n = n.Left
		}
		out = append(out, n.Val)
		for n != nil && n.Right == nil {
			n = pop()
			if n != nil {
				out = append(out, n.Val)
			}
		}
		if n != nil {
			n = n.Right
		}
	}
	return out
}
