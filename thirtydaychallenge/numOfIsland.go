package thirtydaychallenge

var visited [][]bool

func numIslands(grid [][]byte) int {
	var counter = 0
	if len(grid) == 0 {
		return counter
	}
	visited = make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
		for j := range grid[i] {
			visited[i][j] = false
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && visited[i][j] == false {
				traverse(i, j, grid)
				counter++
			}
		}
	}
	return counter
}

func traverse(x int, y int, grid [][]byte) {
	visited[x][y] = true
	if grid[x][y] == '1' {
		if x-1 >= 0 && !visited[x-1][y] {
			traverse(x-1, y, grid)
		}
		if y-1 >= 0 && !visited[x][y-1] {
			traverse(x, y-1, grid)
		}
		if x+1 <= len(grid)-1 && !visited[x+1][y] {
			traverse(x+1, y, grid)
		}
		if y+1 <= len(grid[x])-1 && !visited[x][y+1] {
			traverse(x, y+1, grid)
		}
	}
}
