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

/*
The dependency direction
When you go left to right, 
at index i you'd need to know the max of everything to the 
right — but you haven't seen it yet.
When you go right to left, by the time you reach index i, 
you've already visited every element to its right. 
So curMax has naturally accumulated the answer for i.

*/
