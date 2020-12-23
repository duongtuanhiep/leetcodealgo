package dfs

/*
Idea:
- If we can compute the longest increasing path of every node
we can get the maximum increasing path of the whole matrix.
- Create an additional matrix with the same side to store the
longest increasing path from that specific node.
- On sequence DFS we do not need to recalculate but can append
the computed value if the specific node has been calculated.
*/

// Runtime analysis: O(mn), space O(mn)
func longestIncreasingPath(matrix [][]int) int {
	var res int

	// reset
	maxPath := make([][]int, len(matrix))
	for i := range matrix {
		maxPath[i] = make([]int, len(matrix[i]))
		copy(maxPath[i], matrix[i])
	}
	for i := range maxPath {
		for j := range maxPath[i] {
			maxPath[i][j] = 0
		}
	}

	var cur int
	for i := range matrix {
		for j := range matrix[i] {
			cur, maxPath = longestPath(i, j, matrix, maxPath)
			if cur > res {
				res = cur
			}
		}
	}
	return res
}

// returns max path of current node and an updated version of maxPath grid
func longestPath(i, j int, matrix, maxPath [][]int) (int, [][]int) {

	if maxPath[i][j] > 0 {
		return maxPath[i][j], maxPath
	}

	var left, right, up, down int
	cur := matrix[i][j]
	if j > 0 && cur < matrix[i][j-1] { // Left
		left, maxPath = longestPath(i, j-1, matrix, maxPath)
	}
	if j < len(matrix[i])-1 && cur < matrix[i][j+1] { // right
		right, maxPath = longestPath(i, j+1, matrix, maxPath)
	}
	if i > 0 && cur < matrix[i-1][j] { // up
		up, maxPath = longestPath(i-1, j, matrix, maxPath)
	}
	if i < len(matrix)-1 && cur < matrix[i+1][j] { // down
		down, maxPath = longestPath(i+1, j, matrix, maxPath)
	}

	// updated maxPath
	curMax := 1
	if left >= right && left >= up && left >= down {
		curMax += left
	} else if right >= left && right >= up && right >= down {
		curMax += right
	} else if up >= right && up >= left && up >= down {
		curMax += up
	} else {
		curMax += down
	}

	maxPath[i][j] = curMax
	return curMax, maxPath
}
