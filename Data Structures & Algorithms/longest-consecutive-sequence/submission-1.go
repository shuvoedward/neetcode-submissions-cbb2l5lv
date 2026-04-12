func longestConsecutive(nums []int) int {
mp := map[int]bool{}
	maxL := 0

	for _, num := range nums {
		mp[num] = true
	}

	for _, num := range nums {
		curMax := 0

		if mp[num-1] {
			continue
		}

		for mp[num] {
			num++
			curMax++
		}

		if curMax > maxL {
			maxL = curMax
		}
	}

	return maxL
}
