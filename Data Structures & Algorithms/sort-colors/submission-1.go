func sortColors(nums []int) {
   i, l, r := 0, 0, len(nums)-1

	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
			// decrease i because that value can be 2 itself
			// i-- causes to check that same index again
			i--
		}
		i++
	} 
}
