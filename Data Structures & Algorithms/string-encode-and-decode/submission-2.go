type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded strings.Builder
	for _, s := range strs{
		encoded.WriteString(strconv.Itoa(len(s)) + "#" + s) 
	}

	return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
	// 5#Hello5#World
	decoded := []string{}
	
	for i := 0; i < len(encoded);{
		r := i
		for encoded[r] != '#'{
			r++
		}
		length, _ := strconv.Atoi(encoded[i:r])
		i = r + 1
		decoded = append(decoded, encoded[i : i + length])
		i += length
	}

	return decoded
}
