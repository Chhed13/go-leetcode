package _104_Maximum_Depth_of_Binary_Tree

import "fmt"

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

/*
1. BFS
    - queue of nodes
      move by layers of depth. Check size of queue on start of each wave
2. DFS
    - stack of [nodex,depth]
      +1 to depth on push of new nodes, check leaf status on pop
    - recursion, bottom up approach
      md = max(md(left), md(right)) + 1
*/

// recursive
func maxDepth(root *TreeNode) int {
    if root == nil {
        return  0
    }

    left_depth := maxDepth(root.Left)
    right_depth := maxDepth(root.Right)

    return max(left_depth, right_depth) + 1
}


//queue
type Queue []*TreeNode

func (q *Queue) push(node *TreeNode) {
    *q = append(*q, node)
}

func (q *Queue) pull() *TreeNode {
    if len(*q) == 0 {
        return nil
    }

    fmt.Println(len(*q))
    n := (*q)[0]
    *q = (*q)[1:]
    fmt.Println(len(*q))
    return n
}

func (q *Queue) length() int {
    return len(*q)
}


func maxDepth2(root *TreeNode) int {
    if root == nil {
        return  0
    }

    q := Queue{root}
    h := 0
    for q.length() > 0 {
        for size := q.length(); size > 0; size-- {
            n := q.pull();
            if n.Left != nil {
                q.push(n.Left)
            }
            if n.Right != nil {
                q.push(n.Right)
            }
        }
        h++
    }

    return h
}