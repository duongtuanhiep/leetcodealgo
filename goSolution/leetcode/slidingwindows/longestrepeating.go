package slidingwindows

import "fmt"

/*

Question 424: https://leetcode.com/problems/longest-repeating-character-replacement

**/

func characterReplacement(s string, k int) int {
	var res int = k
	i, j := 0, 0
	count := make([]int, 26)
	count[s[i]-'A']++
	for i < len(s) && j < len(s) {

		// extend till fail
		for isValidReplacement(k, count) && j < len(s) {
			if j-i+1 > res {
				res = j - i + 1
			}
			j++
			if j < len(s) {
				runePos := s[j] - 'A'
				count[runePos]++
			}
			// fmt.Println("ee", count, j)
		}

		// fmt.Println("e", i, j, count)

		if j == len(s) {
			return res
		}

		// retract till success
		for !isValidReplacement(k, count) {
			runePos := s[i] - 'A'
			count[runePos]--
			i++
		}

		fmt.Println("r", i, j, count)
	}

	return res
}

func isValidReplacement(k int, count []int) bool {
	maxIndex := 0
	replacementSum := 0
	for i := range count {
		if i != maxIndex {
			replacementSum += count[i]
		}
		if count[i] > count[maxIndex] {
			replacementSum += count[maxIndex] - count[i]
			maxIndex = i
		}
	}

	// fmt.Println("re", count, replacementSum, k)
	return replacementSum <= k
}

// A much more elegant solution below
func characterReplacement(s string, k int) int {
	l, maxLen := 0, 0
	freq := [26]int{}
	maxFreq := 0

	for r := 0; r < len(s); r++ {
		freq[s[r]-'A']++
		maxFreq = max(maxFreq, freq[s[r]-'A'])
		curLen := r - l + 1

		if curLen-maxFreq > k {
			freq[s[l]-'A']--
			curLen--
			l++
		}

		if maxLen < curLen {
			maxLen = curLen
		}
	}
	return maxLen
}
