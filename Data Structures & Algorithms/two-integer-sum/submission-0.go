func twoSum(nums []int, target int) []int {
   m := map[int]int{}

	for i, num := range nums {
		remainder := target - num
		if index, exists := m[remainder]; exists {
			return []int{index, i}
		}
		m[num] = i
	}

	return []int{}
}
