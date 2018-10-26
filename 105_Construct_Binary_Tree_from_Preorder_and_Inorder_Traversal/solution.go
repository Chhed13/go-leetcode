package _105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

/*https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func getIndex(arr []int, val int) int {
    if arr == nil {
        return -1
    }

    for k, v := range arr {
        if v == val {
            return k
        }
    }
    return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if preorder == nil || inorder == nil || len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := &TreeNode {preorder[0], nil, nil}
    i := getIndex(inorder, root.Val)
    if i < 0 {
        return nil
    }

    root.Left = buildTree(preorder[1:i+1], inorder[:i])
    root.Right = buildTree(preorder[i+1:], inorder[i+1:])

    return root
}
