func canFinish(numCourses int, prerequisites [][]int) bool {
   visited := make(map[int]bool)

	preMap := make(map[int][]int)
	for i := range numCourses {
		preMap[i] = []int{}
	}

	for _, node := range prerequisites {
		crs, pre := node[0], node[1]
		preMap[crs] = append(preMap[crs], pre)
	}

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if visited[crs] {
			return false
		}

		if len(preMap[crs]) == 0 {
			return true
		}

		visited[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}

		visited[crs] = false
		preMap[crs] = []int{}
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true 
}
