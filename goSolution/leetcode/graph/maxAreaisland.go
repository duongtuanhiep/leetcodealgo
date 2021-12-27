package graph

/**
Question 695: https://leetcode.com/problems/max-area-of-island/


Idea:
- BFS/DFS and count, keep track of visited position

Optimization:
- mark visited spot as -1 instead of creating a visit area
*/

func maxAreaOfIsland(grid [][]int) int {
	var visit [][]bool
	for _ = range grid {
		visit = append(visit, make([]bool, len(grid[0])))
	}

	var res int
	for i := range grid {
		for j := range grid[i] {
			if !visit[i][j] {
				if cur := search(grid, visit, i, j); cur > res {
					res = cur
				}
			}
		}
	}

	return res
}

func search(grid [][]int, visited [][]bool, i, j int) int {
	visited[i][j] = true
	if grid[i][j] == 0 {
		return 0
	}

	area := 0
	q := [][]int{[]int{i, j}}
	for len(q) > 0 {
		i, j = q[0][0], q[0][1]
		if i-1 >= 0 && !visited[i-1][j] && grid[i-1][j] == 1 {
			visited[i-1][j] = true
			q = append(q, []int{i - 1, j})
		}
		if j-1 >= 0 && !visited[i][j-1] && grid[i][j-1] == 1 {
			visited[i][j-1] = true
			q = append(q, []int{i, j - 1})
		}
		if i+1 < len(grid) && !visited[i+1][j] && grid[i+1][j] == 1 {
			visited[i+1][j] = true
			q = append(q, []int{i + 1, j})
		}
		if j+1 < len(grid[0]) && !visited[i][j+1] && grid[i][j+1] == 1 {
			visited[i][j+1] = true
			q = append(q, []int{i, j + 1})
		}

		q = q[1:]
		area++
	}
	return area
}
