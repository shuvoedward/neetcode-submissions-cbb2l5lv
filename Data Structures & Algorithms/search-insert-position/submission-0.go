func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r{
		mid := l + (r - l)/2
		if target == nums[mid]{
			return mid
		}
		
		if target > nums[mid]{
			l = mid + 1
		}else if target < nums[mid]{
			r = mid - 1
		}
		
	}

	return l
}
