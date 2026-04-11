func isAnagram(s string, t string) bool {
if len(s) != len(t) {
		return false
	}

	arrS := make([]byte, 26)
	arrT := make([]byte, 26)

	for i := 0; i < len(s); i++ {
		arrS[s[i]-'a']++
		arrT[t[i]-'a']++
	}

	for i := range 26 {
		if arrS[i] != arrT[i] {
			return false
		}
	}

	return true
}
