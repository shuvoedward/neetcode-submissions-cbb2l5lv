func maxAscendingSum(nums []int) int {
   maxSum := nums[0]
	curSum := nums[0]
	
	for i := 1; i < len(nums); i++{
		if nums[i-1] < nums[i]{
			curSum += nums[i]
		}else{
			curSum = nums[i]
		}

		maxSum = max(maxSum, curSum)
	}
	return maxSum
}