package dynamicp

/*
Question 63: https://leetcode.com/problems/unique-paths-ii/

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
- (1): blockages

[-4, -2, (1), -1]
[-2, -2, -2, -1]
[0 , (1), -1, -1]

**/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] = -1
	for i := len(obstacleGrid) - 1; i >= 0; i-- {
		for j := len(obstacleGrid[0]) - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if j+1 < len(obstacleGrid[0]) && obstacleGrid[i][j+1] != 1 {
				obstacleGrid[i][j] += obstacleGrid[i][j+1]
			}
			if i+1 < len(obstacleGrid) && obstacleGrid[i+1][j] != 1 {
				obstacleGrid[i][j] += obstacleGrid[i+1][j]
			}
		}
	}
	return obstacleGrid[0][0] * -1
}
