func getConcatenation(nums []int) []int {
   // i = 0, 
   // 4 + 0 = 4
   // i = 3, 4 + 3= 7
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
