type Solution struct{}

func (s *Solution) Encode(strs []string) string {
var result strings.Builder

	for _, s := range strs {
		result.WriteString(strconv.Itoa(len(s)))
		result.WriteString("*")
		result.WriteString(s)
	}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
res := []string{}
	i := 0

	for i < len(encoded) {
		j := i
		for encoded[j] != '*' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:length+i])
		i += length
	}

	return res
}
