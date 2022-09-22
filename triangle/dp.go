func minimumTotal(triangle [][]int) int {
    
    for row := len(triangle)-2; row > -1; row-- {
        for col := range triangle[row] {
            tmp := triangle[row+1][col]
            if triangle[row+1][col+1] < tmp { tmp = triangle[row+1][col+1] }
            triangle[row][col] += tmp
        }
    }
    
    return triangle[0][0]
}
