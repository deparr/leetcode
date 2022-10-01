func setZeroes(matrix [][]int)  {
    type coord struct {
        row int
        col int
    }
    cache := map[coord]bool {}
    for i, row := range matrix {
        for j, val := range row {
            _, is_new := cache[coord{i, j}]
            if val == 0  && !is_new {
                for tmp := 0; tmp < len(row); tmp++ {
                    if row[tmp] == 0 { continue; }
                    row[tmp] = 0
                    cache[coord{i, tmp}] = true
                }
                for tmp := 0; tmp < len(matrix); tmp++ {
                    if matrix[tmp][j] == 0 { continue; }
                    matrix[tmp][j] = 0
                    cache[coord{tmp, j}] = true
                }
            }
        }
    }
}
