func sortColors(nums []int) {
   i, l, r := 0, 0, len(nums)-1

	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
			i--
		}
		i++
	} 
}
