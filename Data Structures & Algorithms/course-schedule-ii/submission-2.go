func findOrder(numCourses int, prerequisites [][]int) []int {
    prereq := make(map[int][]int)
	visited := make(map[int]bool)
	cycle := make(map[int]bool)

	for i := range numCourses{
		prereq[i] = []int{}
	}

	for _, node := range prerequisites{
		crs, pre := node[0], node[1]
		prereq[crs] = append(prereq[crs], pre)
	}

	output := []int{}
	var dfs func(i int)bool

	dfs = func(i int) bool{
		if cycle[i]{
			return false
		}

		if visited[i]{
			return true
		}

		cycle[i] = true
		for _, crs := range prereq[i]{
			if !dfs(crs){
				return false
			}
		}

		cycle[i] = false
		visited[i] = true
		output = append(output, i)
		return true
	}

	for i := range numCourses{
		if !dfs(i){
			return []int{}
		}
	}

	return output
}
