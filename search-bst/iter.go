/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    t := root
    for t != nil {
        if t.Val == val {
            break
        }

        if val < t.Val {
            t = t.Left
        } else {
            t = t.Right
        }
    }
    return t
}

