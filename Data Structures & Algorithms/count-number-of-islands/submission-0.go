func numIslands(grid [][]byte) int {
   	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	rows, cols := len(grid), len(grid[0])
	islands := 0
	

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= rows ||
			j >= cols || grid[i][j] == '0'  {
			return
		}
		grid[i][j] = '0'
		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1])
		}

	}
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				dfs(r, c)
				islands++
			}
		}
	}
	return islands 
}
