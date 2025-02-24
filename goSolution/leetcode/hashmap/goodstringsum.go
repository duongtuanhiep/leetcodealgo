package hashmap

/*
Question 1160: https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
*/

func countCharacters(words []string, chars string) int {
	char := make(map[byte]int)
	for i := range chars {
		char[chars[i]]++
	}

	var res int
	for _, word := range words {
		wordMap := make(map[byte]int)
		valid := true
		for i := range word {
			wordMap[word[i]]++
			if wordMap[word[i]] > char[word[i]] {
				valid = false
				break
			}
		}
		if valid {
			res += len(word)
		}
	}
	return res
}
