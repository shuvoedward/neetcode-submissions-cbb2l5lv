func characterReplacement(s string, k int) int {
	counts := make([]int, 26)
 	maxLen := 0
	maxFreq := 0
	
	l := 0
	for r := 0; r < len(s); r++ {
		counts[s[r]-'A']++

		maxFreq = max(maxFreq, counts[s[r]-'A'])

		for (r-l+1)- maxFreq > k {
			counts[s[l]-'A']--
			l++
		}

		maxLen = max(maxLen, r-l+1)
	}

	return maxLen	
}
