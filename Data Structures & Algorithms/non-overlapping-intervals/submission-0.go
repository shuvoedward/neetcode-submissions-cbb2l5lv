import (
	"slices"
	"cmp"
)
func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a[]int, b[]int)int{
		return cmp.Compare(a[0], b[0])
	})	
    
	removed := 0
	runningEnd := intervals[0][1]

	for _, interval := range intervals[1:]{
		start, end := interval[0], interval[1]
		if runningEnd > start{
			removed++
			runningEnd = min(end, runningEnd)
		}else{
			runningEnd = end
		}
	}

	return removed
}
