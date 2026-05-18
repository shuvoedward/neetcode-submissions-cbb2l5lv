import "slices"
func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	// 1, 2, 4, 5
	l , r := 0, len(people)-1
	res := 0
	for l <= r{
		remainder := limit - people[r]
		res++
		r--
		if l <= r && people[l] <= remainder{
			l++
		}
	}

	return res

}
