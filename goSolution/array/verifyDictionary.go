package array

/*
Question 953: https://leetcode.com/problems/verifying-an-alien-dictionary/


**/

func isAlienSorted(words []string, order string) bool {
	// Create order map
	orderMap := make(map[byte]int)
	posCounter := 1
	for i := range order {
		orderMap[order[i]] = posCounter
		posCounter++
	}

	for i := range words {
		if i < len(words)-1 {
			if !compareWords(words[i], words[i+1], orderMap) {
				return false
			}
		}
	}

	return true
}

func compareWords(first, second string, order map[byte]int) bool {
	for i := range first {
		if i == len(second) {
			return false
		}
		if order[first[i]] < order[second[i]] {
			return true
		} else if order[first[i]] > order[second[i]] {
			return false
		}
	}
	return true
}
