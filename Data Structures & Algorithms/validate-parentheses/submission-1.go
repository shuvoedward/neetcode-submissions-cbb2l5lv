func isValid(s string) bool {
	stack := []rune{}

	m := map[rune]rune{
		']':'[',
		'}':'{',
		')':'(',
	}

	for _, p := range s{
		if val, exists := m[p]; exists{
			if len(stack) == 0{
				return false
			}
			op := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if op != val{
				return false
			}
		}else{
			stack = append(stack, p)
		}
	}

	return len(stack) == 0
}
