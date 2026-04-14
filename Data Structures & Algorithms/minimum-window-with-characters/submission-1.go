func minWindow(s string, t string) string {
	if t == ""{
		return ""
	}


	countT := map[rune]int{}
	for _, ch := range t {
		countT[ch]++
	}

   	have, need := 0, len(countT)
	resLen := math.MaxInt
	res := [2]int{-1, -1}

	l := 0
	window := map[rune]int{}
	for r, ch := range s {
		window[ch]++

		if countT[ch] > 0 && window[ch] == countT[ch] {
			have++
		}

		for have == need {
			if r-l+1 < resLen {
				resLen = r - l + 1
				res = [2]int{l, r}
			}

			window[rune(s[l])]--
			if countT[rune(s[l])] > 0 && countT[rune(s[l])] > window[rune(s[l])] {
				have--
			}
			l++
		}

	}

	if res[0] == -1 {
		return ""
	}

	return s[res[0] : res[1]+1]
}
