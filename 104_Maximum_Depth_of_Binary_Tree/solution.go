package _104_Maximum_Depth_of_Binary_Tree

/*https://leetcode.com/problems/maximum-depth-of-binary-tree/

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func max(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}

// bottom-up
func maxDepth(root *TreeNode) int {
    if root == nil {
        return  0
    }

    left_depth := maxDepth(root.Left)
    right_depth := maxDepth(root.Right)

    return max(left_depth, right_depth) + 1
}
