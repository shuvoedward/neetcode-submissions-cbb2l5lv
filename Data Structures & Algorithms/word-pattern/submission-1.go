func wordPattern(pattern string, s string) bool {
   strSlc :=  strings.Fields(s) 
   if len(pattern) != len(strSlc){
	return false
   }
	m := map[rune]string{}

	for i, c := range pattern{
		if word, exists := m[c]; exists{
			if word != strSlc[i]{
				return false
			}
		}else{
			for _, val := range m{
				if val == strSlc[i]{
					return false
				}
			}
			m[c] = strSlc[i]
		}
	}

	return true
}