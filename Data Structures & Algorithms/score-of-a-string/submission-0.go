func scoreOfString(s string) int {
	res := 0
	i := 0
	for i < len(s)-1{
		res += abs(int(s[i])-int(s[i+1]))
		i++
	}

	return res
}

func abs(num int)int{
	if num < 0{
		return -num
	}
	return num
}