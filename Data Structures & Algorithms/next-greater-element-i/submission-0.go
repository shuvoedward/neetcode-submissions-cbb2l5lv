func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1Idx := make(map[int]int)
	for i, num := range nums1{
		nums1Idx[num] = i
	}

	res := make([]int, len(nums1))
	for i := range res{
		res[i] = -1
	}

	stack := []int{}
	for _, num := range nums2{
		for len(stack) > 0 && num > stack[len(stack)-1]{
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			idx := nums1Idx[val]
			res[idx] = num
		}
		if _, ok := nums1Idx[num]; ok{
			stack = append(stack, num)
		}
	}
	
	return res
}
