func moveZeroes(nums []int) {
	left := 0
	for right, num := range nums{
		if num != 0{
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}	
}
