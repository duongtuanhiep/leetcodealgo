package stringmanipulation

func isAnagram(s string, t string) bool {
	sCount := make(map[rune]int)

	for _, val := range s {
		sCount[val]++
	}

	for _, val := range t {
		sCount[val]--
	}

	for _, val := range sCount {
		if val != 0 {
			return false
		}
	}

	return true
}
