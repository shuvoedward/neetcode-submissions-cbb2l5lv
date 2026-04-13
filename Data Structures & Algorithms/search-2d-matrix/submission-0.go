func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	l, r := 0, r*c-1

	for l <= r{
		m := l + (r - l)/2
		row, col := m/c, m%c
		if matrix[row][col] < target{
			l = m + 1
		}else if matrix[row][col] > target{
			r = m - 1
		}else{
			return true
		}
	}

	return false
}
