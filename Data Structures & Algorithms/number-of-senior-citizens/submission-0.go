func countSeniors(details []string) int {
   result := 0
   for _, detail := range details{
	if detail[11:13] > "60"{
		result++
	}
   } 

   return result
}