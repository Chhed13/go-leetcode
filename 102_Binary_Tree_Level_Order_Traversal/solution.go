package _102_Binary_Tree_Level_Order_Traversal

/*https://leetcode.com/problems/binary-tree-level-order-traversal/

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) put(t *TreeNode) {
    *q = append((*q), t)
}

func (q *Queue) pull() *TreeNode {
    if len(*q) <= 0 {
        return nil
    }
    t := (*q)[0]
    *q = (*q)[1:]
    return t
}

func levelOrder(root *TreeNode) [][]int {
    qc, qn := Queue{}, Queue{}

    qc.put(root)
    result := [][]int{}
    l := []int{}
    for t := qc.pull(); t != nil; t = qc.pull() {
        l = append(l, t.Val)
        if t.Left != nil {
            qn.put(t.Left)
        }
        if t.Right != nil {
            qn.put(t.Right)
        }

        if len(qc) <= 0 {
            result = append(result, l)
            l = []int{}
            if len(qn) > 0 {
                qc, qn = qn, qc
            } else {
                break
            }
        }
    }

    return result
}
