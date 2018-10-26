package _106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

/*https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
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

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }

    root := TreeNode{postorder[len(postorder)-1], nil, nil}
    i := getIndex(inorder, root.Val)
    if i < 0 {
        return nil
    }

    root.Left = buildTree(inorder[:i],postorder[:i])
    root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

    return &root
}
