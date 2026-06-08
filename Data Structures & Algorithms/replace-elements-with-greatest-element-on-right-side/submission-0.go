func replaceElements(arr []int) []int {
	n := len(arr)
	curMax := -1
	
	for i := n-1; i >=0; i--{
		tmp := arr[i]
		arr[i] = curMax
		curMax = max(curMax, tmp) 
	}

	return arr
}

