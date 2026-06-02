func compress(chars []byte) int {
	n := len(chars)
	k, i := 0, 0

	for i < n{
		chars[k] = chars[i]
		k++

		j := i+1
		for j < n && chars[j] == chars[i]{
			j++
		}

		if j - i > 1{
			cnt := strconv.Itoa(j-i)
			for _, c := range cnt{
				chars[k] = byte(c)
				k++
			}
		}

		i = j
	}

	return k
}
