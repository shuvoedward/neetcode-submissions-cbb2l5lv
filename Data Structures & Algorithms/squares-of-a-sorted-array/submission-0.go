func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	n := len(nums)-1
	left, right := 0, len(nums)-1


	for left <= right{
		lSqr := abs(nums[left]) * abs(nums[left])
		rSqr := abs(nums[right]) * abs(nums[right])

		if rSqr > lSqr{
			result[n] = rSqr	
			right--
		}else{
			result[n] = lSqr
			left++
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
