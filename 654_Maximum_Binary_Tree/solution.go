package _654_Maximum_Binary_Tree

/*
https://leetcode.com/problems/maximum-binary-tree/description/

Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

The root is the maximum number in the array.
The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.
Construct the maximum tree by the given array and output the root node of this tree.

Example 1:
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
Note:
The size of the given array will be in the range [1,1000].
 */

/**
* Definition for a binary tree node.
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	var maxk, maxv, nk, nv int
	for nk, nv = range nums {
		if nv > maxv {
			maxk = nk
			maxv = nv
		}
	}
	maxNode := TreeNode {
		Val: maxv,
		Left: constructMaximumBinaryTree(nums[:maxk]),
		Right: constructMaximumBinaryTree(nums[maxk+1:]),
	}
	return &maxNode
}