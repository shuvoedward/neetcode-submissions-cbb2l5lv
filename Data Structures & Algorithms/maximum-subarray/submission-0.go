func maxSubArray(nums []int) int {
	res := nums[0]
	curSum := 0
   for _, num := range nums{
	if curSum < 0{
		curSum = 0
	}
	curSum += num
	res = max(res, curSum)
   } 
   return res
}
