func majorityElement(nums []int) int {
   candidate, count := 0, 0

   for _, num := range nums{
	if count == 0{
		candidate = num
		count++
	}

	if candidate == num{
		count++
	}else{
		count--
	}
   }

   return candidate
}
