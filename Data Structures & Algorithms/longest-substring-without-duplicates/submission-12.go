func lengthOfLongestSubstring(s string) int {
	
	res, l := 0, 0
	seen := map[byte]int{}

	for r := range s {
		if index, exists := seen[s[r]]; exists && index >= l {
			l = index + 1
		}

		seen[s[r]] = r

		res = max(res, r-l+1)

	}

	return res

}
