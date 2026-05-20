func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	type t struct {
		temperature int
		index       int
	}
	stack := []t{}

	for i, temp := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].temperature < temp {
			lastTemp := stack[len(stack)-1]
			res[lastTemp.index] = i - lastTemp.index
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, t{temperature: temp, index: i})

	}

	return res
}

/*
This is the monotonic decreasing stack 
pattern.The core insight is this: 
you want to maintain a stack of unresolved days
 — days that haven't yet found a warmer future day. 
 The moment a warmer day arrives, it resolves everything 
 colder than it on the stack.
*/
