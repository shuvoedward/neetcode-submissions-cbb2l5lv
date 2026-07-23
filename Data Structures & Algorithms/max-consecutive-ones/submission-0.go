func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0

	for _, num := range nums{
		if num == 0{
			count = 0
		}else{
			count++
		}

		maxCount = max(count, maxCount)
	}	

	return maxCount
}
