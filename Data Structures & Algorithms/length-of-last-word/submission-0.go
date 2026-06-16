func lengthOfLastWord(s string) int {
	newS := strings.Fields(s)
	return len(newS[len(newS)-1])
}
