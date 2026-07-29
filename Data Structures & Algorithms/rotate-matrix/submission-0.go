func rotate(matrix [][]int)  {
   l, r := 0, len(matrix)-1

   for l < r{
	for i := range r - l{
		top, bottom := l, r
		topLeft := matrix[top][l+i]

		matrix[top][l+i] = matrix[bottom-i][l]
		matrix[bottom-i][l] = matrix[bottom][r-i]
		matrix[bottom][r-i] = matrix[top+i][r]
		matrix[top+i][r] = topLeft
	}
	l++
	r--
   } 
}
