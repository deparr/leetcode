/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    order := make([]int, 0, 25)
    inorderhelper(root, &order)
    return order
}

func inorderhelper(node *TreeNode, order *[]int) {
    if node == nil {
        return
    }
    
    inorderhelper(node.Left, order);
    *order = append(*order, node.Val)
    inorderhelper(node.Right, order)
}
