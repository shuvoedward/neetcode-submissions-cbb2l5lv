func decodeString(s string) string {
	stringStack := []string{}
	countStack := []int{}

	cur := ""
	k := 0

	for _, c := range s{
		if c >= '0' && c <= '9'{
			k = k*10 + int(c - '0')
		}else if c == '['{
			stringStack = append(stringStack, cur)
			countStack = append(countStack, k)
			k = 0
			cur = ""
		}else if c == ']'{
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			temp:= cur
			cur = stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			for range count{
				cur+=temp
			}
		}else{
			cur += string(c)
		}
	}

	return cur
}
