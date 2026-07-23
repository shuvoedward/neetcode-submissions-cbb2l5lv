func maxAscendingSum(nums []int) int {
   maxSum := nums[0]
	curSum := nums[0]
	
	// if ascending order, keep adding
	// if not, then start sum from there or reset the old sum
	for i := 1; i < len(nums); i++{
		if nums[i] <= nums[i-1]{
			curSum = 0
		}
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}