func maximumBeauty(items [][]int, queries []int) []int {
    sort.Slice(items, func (x, y int) bool {
        if items[x][1] == items[y][1] {
            return items[x][0] < items[y][0]
        }
        return items[x][1] > items[y][1]
    })

    res := make([]int, len(queries))
    for j, q := range queries {
        cur_max := 0
        for _, val := range items {
            if val[0] <= q {
                cur_max = val[1]
                break
            }
        }
        res[j] = cur_max
    }

    return res
}
