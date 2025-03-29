package dynamicp

/*

Question 1092: https://leetcode.com/problems/shortest-common-supersequence/

Observation:
- This is very similar to the "DNA matching" algorithm (Needleman–Wunsch algorithm)
- What we can do is try to merge the str1 and str2 together so that len(merged) is smallest.
- From the dna matching matrix, we can possibly reconstruct the string somehow

Solution:
We create an 2d arrays matrix for matching algorithm with the following rules:
- Match = 0; mismatch = 1,
- cell values = min(left, top, top-left) + match

From that we go from bottom up and construct the string with rules:
- From current pos matrix[i][j], try to go min(left, top, top-left)
- If it's a match -> go top-left, decrease i,j by one
- If it's not a match, if left <= top, go left, decrease i by one. Else go top, decrease j by one
- Stop when reached blank space

**/

func shortestCommonSupersequence(str1 string, str2 string) string {
	// First, find the longest common subsequence (LCS)
	m, n := len(str1), len(str2)

	// dp[i][j] = length of LCS of str1[0:i] and str2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Construct the shortest common supersequence
	// Start from the bottom right of the dp table
	i, j := m, n
	var result []byte

	for i > 0 && j > 0 {
		// If current characters are same, add it once to result
		if str1[i-1] == str2[j-1] {
			result = append(result, str1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			// Current character of str1 is not part of LCS
			result = append(result, str1[i-1])
			i--
		} else {
			// Current character of str2 is not part of LCS
			result = append(result, str2[j-1])
			j--
		}
	}

	// Add remaining characters
	for i > 0 {
		result = append(result, str1[i-1])
		i--
	}

	for j > 0 {
		result = append(result, str2[j-1])
		j--
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
