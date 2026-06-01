func totalFruit(fruits []int) int {
	count := make(map[int]int)
	l := 0

	for r := range len(fruits){
		count[fruits[r]]++

		if len(count) > 2{
			count[fruits[l]]--
			if count[fruits[l]] == 0{
				delete(count, fruits[l])
			}
			l++
		}
	}

	return len(fruits) - l
}

