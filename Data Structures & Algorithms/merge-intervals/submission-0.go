import (
	"slices" 
	"cmp"
	)
func merge(intervals [][]int) [][]int {
   slices.SortFunc(intervals, func(first []int, second []int) int {
		return cmp.Compare(first[0], second[0])
	})

	res := [][]int{}
	n := len(intervals)

	for i := 0; i < n; {
		cur := intervals[i]
		for i < n-1 && cur[1] >= intervals[i+1][0] {
			cur[0] = min(cur[0], intervals[i+1][0])
			cur[1] = max(cur[1], intervals[i+1][1])
			i++
		}
		res = append(res, cur)
		i++
	}

	return res
}
