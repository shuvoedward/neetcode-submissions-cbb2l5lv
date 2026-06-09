func appendCharacters(s string, t string) int {
	si, ti := 0, 0
	for si < len(s) && ti < len(t){
		if s[si] == t[ti]{
			si++
			ti++
		}else{
			si++
		}
	}

	return len(t) - ti
	
}