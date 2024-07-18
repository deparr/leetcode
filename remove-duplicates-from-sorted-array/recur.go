type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
    sum := 0

    var dfs func(*TreeNode, int64, int)
    dfs = func(cur *TreeNode, path int64, depth int) {
        if cur == nil {
            return
        }

        if cur.Right == nil && cur.Left == nil {
            pathSum := cur.Val
            pow := 10
            for i := depth - 1; i > -1; i-- {
                digit := int((path & (0xf << (i * 4))) >> (i * 4))
                pathSum += digit * pow
                pow *= 10
            }
            sum += pathSum
            return 
        }

        dfs(cur.Left, path | (int64(cur.Val) << (depth * 4)), depth + 1)
        dfs(cur.Right, path | (int64(cur.Val) << (depth * 4)), depth + 1)

    }

    dfs(root, 0, 0)
    return sum
}
