type Node struct {
    Val int
    Children []*Node
}

func levelOrder(root *Node) [][]int {
    levels := [][]int{}
    if root == nil {
        return levels
    }

    queue := []*Node{}
    queue = append(queue, root)

    for len(queue) > 0 {
        level := []int{}
        endOfLevel := len(queue)

        for i := 0; i < endOfLevel; i++ {
            n := queue[i]
            level = append(level, n.Val)

            for _, c := range n.Children {
                queue = append(queue, c)
            }
        }

        levels = append(levels, level)
        queue = queue[endOfLevel:]
    }

    return levels
}
