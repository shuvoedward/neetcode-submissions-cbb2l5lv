func lengthOfLongestSubstring(s string) int {
if len(s) == 0 {
		return 0
	}

	seen := map[rune]int{}

	l := 0
	maxLen := 0

	for r, ch := range s {
		if index, exists := seen[ch]; exists {
			if l <= index {
				l = index + 1
			}
		}

		seen[ch] = r

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}
