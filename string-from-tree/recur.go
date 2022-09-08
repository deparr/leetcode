/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
    "strings"
    "strconv"
)

func tree2str(root *TreeNode) string {
    builder := strings.Builder{}
    helper(root, &builder)
    return builder.String()
}

func helper(node *TreeNode, str *strings.Builder) {
    if node == nil {
        return
    }

    str.WriteString(strconv.Itoa(node.Val))
    if node.Left != nil || node.Right != nil {
        str.WriteString("(")
        helper(node.Left, str)
        str.WriteString(")")
    }
    if node.Right != nil {
        str.WriteString("(")
        helper(node.Right, str)
        str.WriteString(")") 
    }
}
