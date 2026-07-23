func maxAscendingSum(nums []int) int {
   maxSum := nums[0]
	curSum := nums[0]
	
	for i := 1; i < len(nums); i++{
		if nums[i] <= nums[i-1]{
			curSum = 0
		}
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}