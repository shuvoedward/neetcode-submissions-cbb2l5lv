func subarraySum(nums []int, k int) int {
	curSum, res := 0, 0
	prefixSum  := map[int]int{0:1}

	for _, num := range nums{
		curSum += num 
		diff := curSum - k
		
		res+= prefixSum[diff]
		prefixSum[curSum]++
	}
	
	return res
}
