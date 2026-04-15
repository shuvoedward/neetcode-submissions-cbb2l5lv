func maxSlidingWindow(nums []int, k int) []int {
   	result := []int{}

	q := []int{}

	l := 0

	for r := range len(nums) {
		for len(q) > 0 && nums[r] > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, r)

		if l > q[0] {
			q = q[1:]
		}

		if r+1 >= k {
			result = append(result, nums[q[0]])
			l++
		}
	}

	return result 
}
