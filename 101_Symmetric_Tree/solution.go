package _101_Symmetric_Tree

/*https://leetcode.com/problems/symmetric-tree/

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
Note:
Bonus points if you could solve it both recursively and iteratively.
*/

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type nodeTree []*TreeNode

func (n *nodeTree) pull() *TreeNode {
    if len(*n) == 0 {
        return nil
    }

    p := (*n)[0]
    *n = (*n)[1:]
    return p
}

func (n *nodeTree) push(node *TreeNode) {
    *n = append(*n,node)
}

func isSymmetric2(root *TreeNode) bool {
    left := nodeTree{root}
    right := nodeTree{root}
    for ; len(left) > 0 && len(right) > 0;  {
        l,r := left.pull(),right.pull()
        if (l == nil && r == nil) {
            continue
        }
        if (l == nil && r != nil || r == nil && l != nil || l.Val != r.Val) {
            return false
        }

        left.push(l.Left)
        left.push(l.Right)
        right.push(r.Right)
        right.push(r.Left)
    }

    if len(left) > 0 || len(right) > 0 {
        return false
    }
    return true
}

func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return  false
    }

    return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}