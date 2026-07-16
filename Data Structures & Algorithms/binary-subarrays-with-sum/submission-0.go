func numSubarraysWithSum(nums []int, goal int) int {
	prefixSum, res := 0, 0
	count := make(map[int]int)
	count[0] = 1

	for _, num := range nums{
		prefixSum += num
		res += count[prefixSum-goal]
		count[prefixSum]++
	}	

	return res
}
