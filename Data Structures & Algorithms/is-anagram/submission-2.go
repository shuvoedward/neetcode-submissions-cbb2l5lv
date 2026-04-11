func isAnagram(s string, t string) bool {
if len(s) != len(t) {
		return false
	}

	arr := [26]int{}

	for i := range len(s){
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, val := range arr{
		if val != 0{
			return false
		}
	}

	return true
}
