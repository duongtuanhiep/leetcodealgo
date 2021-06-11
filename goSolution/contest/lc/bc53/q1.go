package bc53

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}

	// length 3
	var res int
	for i := range s[:len(s)-2] {
		good := true
		exist := make(map[byte]bool)
		for j := i; j < i+3 && j < len(s); j++ {
			if !exist[s[j]] {
				exist[s[j]] = true
			} else {
				good = false
				break
			}
		}

		if good {
			// fmt.Println(string(s[i : i+2]))
			res++
		}
	}

	return res
}
