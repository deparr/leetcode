type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
    sums := make([]struct{c int; s int}, 10)
    maxDepth := -1

    var explore func (*TreeNode, int)
    explore = func(cur *TreeNode, depth int) {
        if cur == nil {
            return
        }

        if depth >= len(sums) {
            t := struct{c int; s int}{
                c: 0,
                s: 0,
            }
            sums = append(sums, t)
        }
        
        if depth > maxDepth {
            maxDepth = depth
        }

        sums[depth].s += cur.Val
        sums[depth].c += 1

        explore(cur.Left, depth + 1)
        explore(cur.Right, depth + 1)
    }

    explore(root, 0)

    avgs := make([]float64, len(sums))

    for depth, totals := range sums {
        avgs[depth] = float64(totals.s) / float64(totals.c)
    }

    return avgs[:maxDepth+1]
}
