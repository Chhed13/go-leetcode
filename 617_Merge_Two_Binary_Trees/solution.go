package _617_Merge_Two_Binary_Trees

/*https://leetcode.com/problems/merge-two-binary-trees/

Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:

Input:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
Merged tree:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7


Note: The merging process must start from the root nodes of both trees.
*/

type TreeNode struct {
    Val   int
    Left  TreeNode
    Right TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }

    root := TreeNode{}
    var t1L, t1R, t2L, t2R *TreeNode
    if t1 != nil {
        root.Val += t1.Val
        t1L = t1.Left
        t1R = t1.Right
    }
    if t2 != nil {
        root.Val += t2.Val
        t2L = t2.Left
        t2R = t2.Right
    }

    root.Left = mergeTrees(t1L, t2L)
    root.Right = mergeTrees(t1R, t2R)

    return &root
}
