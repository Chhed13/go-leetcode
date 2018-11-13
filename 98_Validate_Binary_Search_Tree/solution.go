package _98_Validate_Binary_Search_Tree

import "math"

/*https://leetcode.com/problems/validate-binary-search-tree/

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:

Input:
    2
   / \
  1   3
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
             is 5 but its right child's value is 4.
*/


/* Description
1. Make in-order traversal (left-root-right) and put all vals into list. All values in the list must be in incremental order
    - O(n) time, O(n.val + max depth) memory
2. Top down solution
    - DFS (stack, recursion) or BFS (queue) doesnt metter, but we must pass the max|min limits everywhere
    - check rule of root.Val with root.Left | root.Right and max|min limits, if ok - move deeper, if rules ok everywhere - return true
    - O(n) - time, O(n) memory
*/

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type LimNode struct {
    node *TreeNode
    min  int
    max  int
}

type Stack []LimNode

func (s *Stack) push(n LimNode) {
    *s = append(*s, n)
}

func (s *Stack) pop() LimNode {
    if len(*s) == 0 {
        return LimNode{nil, 0, 0}
    }

    n := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return n
}

func (s *Stack) length() int {
    return len(*s)
}

// 2. Top Down solution
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    s := Stack{LimNode{root, math.MinInt64, math.MaxInt64}}

    for s.length() > 0 {
        n := s.pop()
        if n.node.Left != nil {
            // fmt.Printf("left %+v \n",n)
            if n.node.Left.Val < n.node.Val && n.node.Left.Val > n.min {
                s.push(LimNode{n.node.Left, n.min, n.node.Val})
            } else {
                return false
            }
        }
        if n.node.Right != nil {
            // fmt.Printf("right %+v \n",n)
            if n.node.Right.Val > n.node.Val && n.node.Right.Val < n.max {
                s.push(LimNode{n.node.Right, n.node.Val, n.max})
            } else {
                return false
            }
        }
    }
    return true
}
