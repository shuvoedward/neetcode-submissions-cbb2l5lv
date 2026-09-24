func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1

	for r - l >= k{
		if abs(x-arr[l]) <= abs(x-arr[r]){
			r--
		}else{
			l++
		}
	}

	return arr[l: r+1]
}

func abs(a int)int{
	if a < 0{
		return -a
	}
	return a
}