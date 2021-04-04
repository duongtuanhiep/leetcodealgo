package dynamicp

import (
	"sort"
)

/*
Question 1048: https://leetcode.com/problems/longest-string-chain/

finding longest words chain of predecessor

Idea:
- Create a supporting function to check for predecessor.
- Sort all the words by length. The possibility that the desired result chains will start from the words with the shortest length.
- Recursion. MaxStep[i] = 1 || Max(MaxStep[i-n])+1
- Runtime: O(N^2*M) : N is the length of array words, M is maximum size of the words.
**/

func longestStrChain(words []string) int {
	// sorting array of string
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	maxStep := make([]int, len(words))
	max := 1
	for i := range words {
		if maxStep[i] == 0 {
			maxStep[i] = 1
		}
		for pos := i - 1; pos >= 0; pos-- {
			if len(words[pos]) == len(words[i])-1 {
				match := isGood(words[pos], words[i])
				if match {
					if maxStep[pos]+1 > maxStep[i] {
						maxStep[i] = maxStep[pos] + 1
					}
					if maxStep[i] > max {
						max = maxStep[i]
					}
				}
			} else if len(words[pos]) < len(words[i])-1 {
				break
			}
		}
	}

	return max
}

// is good compare 2 string to see if first one is predecessor of second one
func isGood(first, second string) bool {
	valid := true
	misMatch := true
	for i, j := 0, 0; i < len(first) && j < len(second); i, j = i+1, j+1 {
		if first[i] != second[j] && misMatch {
			misMatch = false
			i--
		} else if !misMatch && first[i] != second[j] {
			valid = false
		}
	}
	return valid
}
