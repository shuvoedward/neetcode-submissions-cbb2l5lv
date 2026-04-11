func getConcatenation(nums []int) []int {
   l := len(nums)
   
   if l < 1{
	return []int{}
   }
   
   res := make([]int, 2*l)

   for i, num := range nums{
	res[i] = num
	res[l+i] = num
	}

   return res
}
