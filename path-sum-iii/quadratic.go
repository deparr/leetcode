/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    nsums := 0
    var computesum func(*TreeNode, int)
    computesum = func(node *TreeNode, rsum int) {
        if node.Val + rsum == targetSum {
            fmt.Println(node.Val)
            nsums++
        }


        if node.Left != nil {
            computesum(node.Left, rsum + node.Val)
        }

        if node.Right != nil {
            computesum(node.Right, rsum + node.Val)
        }
    }

    var computeallsums func(*TreeNode)
    computeallsums = func(node *TreeNode) {
        if node == nil {
            return
        }
        computesum(node, 0)
        computeallsums(node.Left)
        computeallsums(node.Right)

    }

    computeallsums(root)

    return nsums
}
