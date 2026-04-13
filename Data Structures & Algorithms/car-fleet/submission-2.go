import "slices"
func carFleet(target int, position []int, speed []int) int {
fleetArr := make([][]int, len(position))

	for i := range len(position) {
		fleetArr[i] = []int{position[i], speed[i]}
	}

	slices.SortFunc(fleetArr, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	stack := []float64{}

	for i := len(fleetArr) - 1; i >= 0; i-- {
		time := float64(target - fleetArr[i][0]) / float64(fleetArr[i][1])
		stack = append(stack, time)
		if len(stack) >=2 && stack[len(stack)-1] <= stack[len(stack)-2]{
			stack = stack[:len(stack)-1]
		}
		
	}

	return len(stack)
}
