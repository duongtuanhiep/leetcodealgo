package dynamicp

/*
Question 63: https://leetcode.com/problems/unique-paths/

Observation: A node with coordinate [i,j] can only moves to [i+1,j] or [i,j+1],
that means that the total number of ways to reach finnish at current node will
be equal to the sum of total number of ways to reach finnish from [i+1,j] and [i,j+1]

Idea: Try to calculate the total number of ways to reach finish from each available
nodes. We can start with the direct neighbor from the destination node.
Instead of using a new grid to calculate, we can use the existing one and start counting
backward from -1. The result will be the absolute value at node[0][0]. The finnish node can
have value -1 (can consider that there is 1 way to reach itself)
S[i][j] = S[i+1][j] + S[i][j+1]

Example:

**/

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := range matrix {
		line := make([]int, n)
		matrix[i] = line
	}

	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}

	return matrix[m-1][n-1]
}
