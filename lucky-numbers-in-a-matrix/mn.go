func luckyNumbers (matrix [][]int) []int {
    lucky := make([]int, 0, 5)

    minRow := make([]int, 0, len(matrix))
    minCol := make([]int, 0, len(matrix[0]))
    for _, row := range matrix {
        minRow = append(minRow, slices.Min(row))
    }

    for x := range len(matrix[0]) {
        col := 0
        for _, row := range matrix {
            if row[x] > col {
                col = row[x]
            }
        }
        minCol = append(minCol, col)
    }

    for _, row := range minRow {
        for _, col := range minCol {
            if row == col {
                lucky = append(lucky, row)
            }
        }
    }

    return lucky
}
