func subarraySum(nums []int, k int) int {
	curSum, res := 0, 0
	prefixSum  := map[int]int{0:1}

	for _, num := range nums{
		curSum += num 
		diff := curSum - k
		
		res+= prefixSum[diff]
		prefixSum[curSum]++
	}
	
	return res
}

/*

Subarray Sum Equals K — Prefix Sum
Instead of asking "does this subarray sum to k?", 
flip the question: "have I seen a prefix sum that, 
if chopped off the front, leaves exactly k?"
If curSum = a + b + c + d, and a + b is some earlier prefix, 
then curSum - (a + b) == k. Rearranged: the prefix I'm looking 
for is curSum - k.
So at every index, compute diff = curSum - k and check the map. 
If that prefix sum was seen before, the elements between that prefix 
and now form a valid subarray summing to k. If it was seen 3 times, 
that's 3 different starting points — all valid.
The map stores how many times each prefix sum has appeared, 
not the subarrays themselves.
The {0: 1} seed handles the case where the subarray 
starts at index 0 — there's no real prefix to subtract, 
so you plant a fake one with value 0, count 1.

*/