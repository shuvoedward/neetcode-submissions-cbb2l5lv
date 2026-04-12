func mergeAlternately(word1 string, word2 string) string {
	l := len(word1)
	
	if len(word2) < l{
		l = len(word2)
	}

	var result strings.Builder
	i := 0
	for i = range l{
		result.WriteByte(word1[i])
		result.WriteByte(word2[i])
	}

	if len(word1) > l{
		result.WriteString(word1[i+1:])
	}
	if len(word2) > l{
		result.WriteString(word2[i+1:])
	}

	return result.String()
}
