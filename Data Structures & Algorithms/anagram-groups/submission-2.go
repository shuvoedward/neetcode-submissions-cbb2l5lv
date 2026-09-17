func groupAnagrams(strs []string) [][]string {
	groupM := map[[26]int][]string{}

	for _, s := range strs {
		var count [26]int

		for _, c := range s {
			count[c-'a']++
		}

		groupM[count] = append(groupM[count], s)
	}

	var result [][]string
	for _, value := range groupM {
		result = append(result, value)
	}

	return result
}

// var count [52]int
// for _, c := range s {
//     if c >= 'a' && c <= 'z' {
//         count[c-'a'+26]++
//     } else if c >= 'A' && c <= 'Z' {
//         count[c-'A']++
//     }
// }