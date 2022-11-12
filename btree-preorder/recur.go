/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    
    res := make([]int, 0, 10)
    var preord func(*TreeNode)
    preord = func(node *TreeNode) {
        res = append(res, node.Val)

        if node.Left != nil {
            preord(node.Left)
        }

        if node.Right != nil {
            preord(node.Right)
        }
    }

    if root != nil {
        preord(root)
    }
    return res
}
