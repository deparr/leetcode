func minFallingPathSum(matrix [][]int) int {
    min := func (a []int) int {
        _min := a[0]
        for _, v := range a[1:] { if v < _min { _min = v }}
        return _min
    }

    for i := len(matrix)-2; i >= 0; i-- {
        for j := 0; j < len(matrix); j++ {
            begin, end := j, j+1
            if j > 0 {
                begin--
            }
            if j < len(matrix)-1 {
                end++
            }
            matrix[i][j] += min(matrix[i+1][begin:end])
        }
    }


    return min(matrix[0])
}
