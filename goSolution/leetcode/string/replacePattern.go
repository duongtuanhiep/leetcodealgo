package string

/*
Question 890: https://leetcode.com/problems/find-and-replace-pattern/

Idea:
- Create a map of swap and a map to check if chars exist
- Go through every sinlge string


**/

func findAndReplacePattern(words []string, pattern string) []string {
	var res []string
	for _, word := range words {
		if len(word) == len(pattern) {
			perm, check, match := make(map[byte]byte), make(map[byte]bool), true
			for i := range pattern {
				if _, ok := perm[word[i]]; !ok && !check[pattern[i]] {
					perm[word[i]] = pattern[i]
					check[pattern[i]] = true
				} else if perm[word[i]] != pattern[i] {
					match = false
					break
				}
			}
			if match {
				res = append(res, word)
			}
		}
	}

	return res
}
