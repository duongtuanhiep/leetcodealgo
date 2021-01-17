package contestOne

func countConsistentStrings(allowed string, words []string) int {
	allow := make(map[rune]bool)
	for _, val := range allowed {
		allow[val] = true
	}

	var count int
	for i := range words {
		for j, char := range words[i] {
			if !allow[char] {
				break
			}
			if j == len(words[i])-1 {
				count++
			}
		}
	}
	return count
}
