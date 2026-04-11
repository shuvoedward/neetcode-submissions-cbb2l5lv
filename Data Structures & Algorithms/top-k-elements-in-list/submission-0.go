func topKFrequent(nums []int, k int) []int {
freq := map[int]int{}
	count := make([][]int, len(nums)+1)

	for _, num := range nums {
		freq[num]++
	}

	for num, cnt := range freq {
		count[cnt] = append(count[cnt], num)
	}

	res := []int{}

	for i := len(count) - 1; i > 0; i-- {
		for _, c := range count[i] {
			res = append(res, c)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
