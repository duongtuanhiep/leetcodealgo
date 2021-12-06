package dynamicp

/*
Question 778: https://leetcode.com/problems/swim-in-rising-water/



testCase :
[[0,2],[1,3]]
[[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
**/

func swimInWater(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid))
	}

	length := len(grid) - 1
	for i := length; i >= 0; i-- {
		for j := length; j >= 0; j-- {
			if i == length && j == length {
				visited[i][j] = true
				continue
			}
			left, right, up, down := 2500, 2500, 2500, 2500
			if i != length && visited[i+1][j] {
				down = grid[i+1][j]
			}
			if i != 0 && visited[i-1][j] {
				up = grid[i-1][j]
			}
			if j != length && visited[i][j+1] {
				right = grid[i][j+1]
			}
			if j != 0 && visited[i][j-1] {
				left = grid[i][j-1]
			}
			grid[i][j] = max(grid[i][j], minFour(left, right, up, down))
			visited[i][j] = true
		}
	}

	return grid[0][0]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minFour(a, b, c, d int) int {
	return min(min(a, b), min(c, d))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
