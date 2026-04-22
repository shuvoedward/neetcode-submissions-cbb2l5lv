func isValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for r := range 9{
		for c := range 9{
			if board[r][c] == '.'{
				continue
			}

			val := board[r][c] - '1'
			bit := 1 << val
			squareIdx := (r/3)*3 + c/3

			if rows[r] & bit != 0 || cols[c]&bit != 0 || 
			squares[squareIdx]&bit != 0{
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			squares[squareIdx] |= bit
		}
	}

	return true
	// n^2 time complexity although each cell is visited once
	// because, area of square is n*n = n^2. so n is not the total number of cells
	// but n is one line length, then 9 * 9 = 81 total 
}
