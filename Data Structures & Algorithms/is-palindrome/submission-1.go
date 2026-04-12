func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	srune := []rune(s)

	for l < r {
		for l < r && !isAlpha(srune[l])  {
			l++
		}
		for l < r && !isAlpha(srune[r]) {
			r--
		}
		if unicode.ToLower(srune[l]) != unicode.ToLower(srune[r]) {
			return false
		}
		l++
		r--
	}

	return true
}
func isAlpha(r rune) bool {
	if unicode.IsDigit(r) || unicode.IsLetter(r) {
		return true
	}
	return false
}