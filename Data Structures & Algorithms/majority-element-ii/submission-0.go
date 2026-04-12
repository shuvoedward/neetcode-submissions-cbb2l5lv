func majorityElement(nums []int) []int {
	count := make(map[int]int)
	for _, num := range nums{
		count[num]++
	}

	res := []int{}
	for key, val := range count{
		if val > len(nums)/3{
			res = append(res, key)
		}
	}

	return res
}
