func calPoints(operations []string) int {
	stack := []int{}

	for _, operation := range operations {
		if operation == "+" {
			if len(stack) < 2 {
				break
			}
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = append(stack, a+b)
		} else if operation == "D" {
			if len(stack) < 1 {
				break
			}
			stack = append(stack, stack[len(stack)-1]*2)
		} else if operation == "C" {
			if len(stack) < 1 {
				break
			}
			stack = stack[:len(stack)-1]
		} else {
			digit, err := strconv.Atoi(operation)
			if err != nil {
				break
			}
			stack = append(stack, digit)
		}

	}

	var result int
	for _, val := range stack {
		result += val
	}

	return result
}
