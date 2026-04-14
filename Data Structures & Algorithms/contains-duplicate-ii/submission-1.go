func containsNearbyDuplicate(nums []int, k int) bool {
	mp := map[int]int{}

	for i, n := range nums{
		if index, exists := mp[n]; exists && i-index<=k{
			return true
		}
		mp[n]=i
	}

	return false
}
