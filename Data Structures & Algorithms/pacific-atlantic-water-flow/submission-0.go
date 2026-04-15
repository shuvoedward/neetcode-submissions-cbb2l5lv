func pacificAtlantic(heights [][]int) [][]int {
   rows, cols := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)

	var dfs func(row, col int, visit map[[2]int]bool, prevHeight int)

	dfs = func(row, col int, visit map[[2]int]bool, prevHeight int) {
		if row < 0 || row >= rows ||
			col < 0 || col >= cols || visit[[2]int{row, col}] ||
			heights[row][col] < prevHeight {
			return
		}

		visit[[2]int{row, col}] = true
		dfs(row+1, col, visit, heights[row][col])
		dfs(row-1, col, visit, heights[row][col])
		dfs(row, col+1, visit, heights[row][col])
		dfs(row, col-1, visit, heights[row][col])
	}

	for c := range cols {
		dfs(0, c, pac, heights[0][c])
		dfs(rows-1, c, atl, heights[rows-1][c])
	}

	for r := range rows {
		dfs(r, 0, pac, heights[r][0])
		dfs(r, cols-1, atl, heights[r][cols-1])
	}

	result := [][]int{}
	for r := range rows {
		for c := range cols {
			node := [2]int{r, c}
			if pac[node] && atl[node] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
