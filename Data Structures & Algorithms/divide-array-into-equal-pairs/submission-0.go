func divideArray(nums []int) bool {
   if len(nums) % 2 != 0{
	return false
   }

   count := map[int]int{}
   pairs := 0

   for _, num := range nums{
	count[num]++
	if count[num] % 2 == 0{
		pairs++
	}
   }

   return pairs == len(nums) / 2
}