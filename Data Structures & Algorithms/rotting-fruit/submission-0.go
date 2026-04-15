func orangesRotting(grid [][]int) int {
   	time := 0
	fresh := 0

	q := [][2]int{}
	rows, cols := len(grid), len(grid[0])

	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 {
				fresh++
			}
			if grid[r][c] == 2 {
				q = append(q, [2]int{r, c})
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for fresh > 0 && len(q) > 0 {
		length := len(q)

		for range length {
			current := q[0]
			q = q[1:]

			for _, dir := range directions {
				newRow := current[0] + dir[0]
				newCol := current[1] + dir[1]

				if newRow >= 0 && newCol >= 0 &&
					newRow < rows && newCol < cols &&
					grid[newRow][newCol] == 1 {
						grid[newRow][newCol] = 2
					q = append(q, [2]int{newRow, newCol})
					fresh--
				}
			}
		}
		time++
	}

	if fresh == 0 {
		return time
	}

	return -1 
}
