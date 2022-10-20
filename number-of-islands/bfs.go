func numIslands(grid [][]byte) int {
    num_islands := 0
    type coord struct {x int; y int}
    visited := map[coord]struct{} {}
    var is_island func (x int, y int)
    is_island = func (x int, y int) {
        _, prs := visited[coord{x, y}]
        if prs { return }
        
        if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' { return }
        
        visited[coord{x, y}] = struct{} {}
        
        is_island(x-1, y)
        is_island(x+1, y)
        is_island(x, y-1)
        is_island(x, y+1)
    }

    for i, row := range grid {
        for j, point := range row {
            _, prs := visited[coord{i, j}]
            if point == '1' &&  !prs{
                num_islands++
                is_island(i, j)
            }
        }
    }
    return num_islands
}

