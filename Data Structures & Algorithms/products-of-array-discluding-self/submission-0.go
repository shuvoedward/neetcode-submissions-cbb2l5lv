func productExceptSelf(nums []int) []int {
	prefix := 1

	result := make([]int, len(nums))

	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	postfix := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
