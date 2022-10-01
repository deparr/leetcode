
/* WORK IN PROGRESS */

func setZeroes(matrix [][]int)  {
    type coord struct {
        row int
        col int
    }
    lower_bound := coord{0, 0}
    for i, row := range matrix {
        for j, val := range row {
            fmt.Println(i, j)
            if val == 0  && i <= lower_bound.row && j >= lower_bound.col {
                found_lb_row, found_lb_col := false, false
                for tmp := j+1; tmp < len(row); tmp++ {
                    if row[tmp] == 0 {
                        found_lb_col = true
                        lower_bound.col = tmp
                        break
                    }
                    row[tmp] = 0
                }
                for tmp := j-1; tmp > -1; tmp-- {
                    if row[tmp] == 0 { break; }
                    row[tmp] = 0
                }
                for tmp := i+1; tmp < len(matrix); tmp++ {
                    if matrix[tmp][j] == 0 {
                        found_lb_col = true
                        lower_bound.row = tmp
                        break
                    }
                    matrix[tmp][j] = 0
                }
                for tmp := i-1; tmp > -1; tmp-- {
                    if matrix[tmp][j] == 0 { break; }
                    matrix[tmp][j] = 0
                }
                if !found_lb_col {
                    lower_bound.col = len(matrix[0]) - 1
                }
                if !found_lb_row {
                    lower_bound.row = len(matrix) - 1
                }
                fmt.Println(lower_bound)
            }
        }
    }
}
