func getConcatenation(nums []int) []int {
   n := len(nums)
   if n == 0{
      return []int{}
   }

   ans := make([]int, n*2)

   for i, num := range nums{
      ans[i] = num
      ans[n+i] = num
   }

   return ans
}
