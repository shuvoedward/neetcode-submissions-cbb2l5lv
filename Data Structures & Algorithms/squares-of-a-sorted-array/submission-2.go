func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	n := len(nums)-1
	left, right := 0, len(nums)-1


	for left <= right{	
		if abs(nums[left]) > abs(nums[right]){
			result[n] = nums[left] * nums[left]	
			left++
		}else{
			result[n] = nums[right] * nums[right]
			right--
		}
		n--
	}
	return result	
}

func abs(x int)int{
	if x < 0{
		return -x
	}
	return x
}
