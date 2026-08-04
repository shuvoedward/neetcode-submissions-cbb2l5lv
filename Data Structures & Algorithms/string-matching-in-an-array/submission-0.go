func stringMatching(words []string) []string {
   res := []string{}

   for i := range words{
	for j := range words{
		if i == j {
			continue		
		}

		if strings.Contains(words[j], words[i]){
			res = append(res, words[i])
			break
		}
	}
   } 
   return res
}