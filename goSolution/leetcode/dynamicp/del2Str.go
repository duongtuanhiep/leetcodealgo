package dynamicp

/*
Question 583: https://leetcode.com/problems/delete-operation-for-two-strings/

Variance of Needleman-Wunsch: https://en.wikipedia.org/wiki/Needleman%E2%80%93Wunsch_algorithm
**/

func minDistance(word1 string, word2 string) int {

	var grid [][]int
	for i := 0; i < len(word2)+1; i++ {
		grid = append(grid, make([]int, len(word1)+1))
	}

	for i := 0; i < len(word2)+1; i++ {
		for j := 0; j < len(word1)+1; j++ {
			if i == 0 {
				grid[i][j] = j
			} else if j == 0 {
				grid[i][j] = i
			} else if word2[i-1] == word1[j-1] {
				grid[i][j] = grid[i-1][j-1]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + 1
			}
		}
	}

	return grid[len(word2)][len(word1)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
