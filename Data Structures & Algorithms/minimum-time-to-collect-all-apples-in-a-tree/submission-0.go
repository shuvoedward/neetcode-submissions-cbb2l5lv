func minTime(n int, edges [][]int, hasApple []bool) int {
	adj := make([][]int, n)
	for i := range adj{
		adj[i] = []int{}
	}

	for _, edge := range edges{
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	var dfs func(cur, parent int)int
	dfs = func(cur, parent int)int{
		time := 0
		for _, child := range adj[cur]{
			if child == parent{
				continue
			}
			childTime := dfs(child, cur)
			if childTime > 0 || hasApple[child]{
				time += 2 + childTime
			}
		}
		return time
	}

	return dfs(0, -1)
}
