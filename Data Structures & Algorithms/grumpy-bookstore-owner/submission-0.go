func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	total := 0
	maxSum := 0

	for i := range minutes {
		if grumpy[i] == 0 {
			total += customers[i]
		} else {
			maxSum += customers[i]
		}
	}
	curSum := maxSum
	for i := minutes; i < len(grumpy); i++ {
		
		if grumpy[i] == 0 {
			total += customers[i]
		} else {
			curSum += customers[i]
		}

		if grumpy[i-minutes] == 1 {
			curSum -= customers[i-minutes]
		}

		maxSum = max(maxSum, curSum)
	}

	return total + maxSum
}
