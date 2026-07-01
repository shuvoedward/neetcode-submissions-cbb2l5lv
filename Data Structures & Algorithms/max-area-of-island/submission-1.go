func maxAreaOfIsland(grid [][]int) int {
   visited := map[[2]int]bool{}
	rows, cols := len(grid), len(grid[0])

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= rows ||
			c >= cols || grid[r][c] == 0 || visited[[2]int{r, c}] {
			return 0
		}
		visited[[2]int{r, c}] = true
		return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
	}

	area := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 {
				area = max(area, dfs(r, c))
			}
		}
	}

	return area
}
