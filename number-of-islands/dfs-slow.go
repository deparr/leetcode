func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
    type coord struct {x int; y int}
    visited := map[coord]struct{} {}
	num_islands := 0

    var dfs func (x int, y int)
    dfs = func (x int, y int) {

		if x < 0 || x >= m { return }
		if y < 0 || y >= n { return }
		_, present := visited[coord{x, y}]
        if present { return }
		if grid[x][y] == '0' { return }
        
        visited[coord{x, y}] = struct{} {}
        
        dfs(x - 1, y)
        dfs(x + 1, y)
        dfs(x, y - 1)
        dfs(x, y + 1)
    }

    for i, row := range grid {
        for j, point := range row {
            _, present := visited[coord{i, j}]
            if point == '1' && !present {
                num_islands++
                dfs(i, j)
            }
        }
    }
    return num_islands
}

