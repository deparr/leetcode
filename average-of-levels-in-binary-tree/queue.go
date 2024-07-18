type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
    queue := []*TreeNode{}
    avgs := []float64{}

    queue = append(queue, root)
    for len(queue) > 0 {
        lenOfLevel := len(queue)
        total := 0.0

        for i := 0; i < lenOfLevel; i++ {
            node := queue[i]

            total += float64(node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        avgs = append(avgs, total / float64(lenOfLevel))
        queue = queue[lenOfLevel:]
    }

    return avgs
}
