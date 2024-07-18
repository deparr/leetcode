type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    levels := [][]int{}
    queue := []*TreeNode{}
    if root != nil {
        queue = append(queue, root)
    }

    for len(queue) > 0 {
        level := []int{}
        levelLen := len(queue)
        for i := 0; i < levelLen; i++ {
            n := queue[i]
            level = append(level, n.Val)

            if n.Left != nil {
                queue = append(queue, n.Left)
            }

            if n.Right != nil {
                queue = append(queue, n.Right)
            }
        }

        levels = append(levels, level)
        queue = queue[levelLen:]
    }

    return levels
}
