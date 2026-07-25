
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, k - 1
	min := math.MaxInt32

	for r < len(nums){
		if nums[r] - nums[l] < min{
			min = nums[r] - nums[l]
		}
		l++
		r++
	}

	return min
}
