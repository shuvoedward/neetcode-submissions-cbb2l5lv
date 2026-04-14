func islandsAndTreasure(grid [][]int) {
   q := [][2]int{}
	rows, cols := len(grid), len(grid[0])
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 0 {
				q = append(q, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		row, col := node[0], node[1]

		for _, dir := range directions {
			r, c := row+dir[0], col+dir[1]

			if r < 0 || c < 0 || r >= rows || c >= cols ||
				grid[r][c] != 2147483647 {
				continue
			}

			q = append(q, [2]int{r, c})
			grid[r][c] = grid[row][col] + 1
		}
	} 
}
