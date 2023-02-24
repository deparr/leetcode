/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	var inv func(*TreeNode)
	inv = func(node *TreeNode) {
		if node == nil {
			return
		}

		node.Left, node.Right = node.Right, node.Left

		inv(node.Left)
		inv(node.Right)
	}

	inv(root)
	return root
}

