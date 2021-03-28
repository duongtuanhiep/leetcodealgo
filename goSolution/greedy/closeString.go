package greedy

import "sort"

/*
Question 1657: https://leetcode.com/problems/determine-if-two-strings-are-close/

Idea:
- Count occurence of each char in word1 and word2
- 2 words are close if for any character 'a1' in word1 with occurence x there will be a character
a2 in word2 with the same occurence.

Optimization:
- Checks the length of word1, word2.
- While creating the bucket for word2 we can also do check to see if there is any character that
only exist in word2.
**/

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	bucket1 := make([]int, 26)
	bucket2 := make([]int, 26)
	for _, char := range word1 {
		bucket1[int(char)-int('a')]++
	}

	for _, char := range word2 {
		if bucket1[int(char)-int('a')] == 0 {
			return false
		}
		bucket2[int(char)-int('a')]++
	}

	sort.Ints(bucket1)
	sort.Ints(bucket2)
	for i := range bucket1 {
		if bucket1[i] != bucket2[i] {
			return false
		}
	}

	return true
}
