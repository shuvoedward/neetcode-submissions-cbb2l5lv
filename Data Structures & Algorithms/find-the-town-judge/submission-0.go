func findJudge(n int, trust [][]int) int {
in := map[int]int{}
	out := map[int]int{}

	for _, t := range trust {
		out[t[0]]++
		in[t[1]]++
	}

	for person, incoming := range in {
		if incoming == n-1 && out[person] == 0 {
			return person
		}
	}

	return -1
}
