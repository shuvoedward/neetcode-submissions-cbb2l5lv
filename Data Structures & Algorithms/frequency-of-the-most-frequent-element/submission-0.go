func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	var total int64 = 0
	res, l := 0, 0

	for r := range len(nums){
		total += int64(nums[r])

		for int64(nums[r])*int64(r-l+1) > total + int64(k){
			total -= int64(nums[l])
			l++
		}

		res = max(res, r-l+1)
	}	

	return res
}
