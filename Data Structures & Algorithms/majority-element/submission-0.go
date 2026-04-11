func majorityElement(nums []int) int {
   count := map[int]int{}

	d := len(nums) / 2

	for _, num := range nums {
		count[num]++
		if count[num] > d {
			return num
		}
	}

	return 0 
}
