func characterReplacement(s string, k int) int {
 	maxLen := 0
	srune := []rune(s)
	l := 0
	m := map[rune]int{}
	maxCh := 0

	for r := 0; r < len(srune); r++ {
		m[srune[r]]++

		if maxCh < m[srune[r]] {
			maxCh = m[srune[r]]
		}

		if r-l+1-maxCh > k {
			m[srune[l]]--
			l++
			
		}

		maxLen = max(maxLen, r-l+1)
	}

	return maxLen	
}
