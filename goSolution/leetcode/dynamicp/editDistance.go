package dynamicp

/*
Question 72: https://leetcode.com/problems/edit-distance/

Idea: Sequence Alignment
*/

func minDistance(word1 string, word2 string) int {
	var match [][]int

	for i := 0; i <= len(word2); i++ {
		match = append(match, make([]int, len(word1)+1))
	}

	for i := range match[0] {
		match[0][i] = i
	}
	for i := range match {
		match[i][0] = i
	}

	for i := 1; i <= len(word2); i++ {
		for j := 1; j <= len(word1); j++ {
			if word1[j-1] == word2[i-1] {
				match[i][j] = match[i-1][j-1]
			} else {
				match[i][j] = min(match[i-1][j], min(match[i-1][j-1], match[i][j-1])) + 1
			}
		}
	}

	// 	for i := range match {
	// 	fmt.Println(match[i])
	// 	}

	return match[len(word2)][len(word1)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
