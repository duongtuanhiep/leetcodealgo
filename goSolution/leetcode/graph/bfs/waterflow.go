package bfs

/*

Question 417: https://leetcode.com/problems/pacific-atlantic-water-flow/

O(N^2) solutions:
- On each cells, try to traverse as far as possible to see if it reach both topleft and bottomright

O(N) solutions: Recursion.
- From each unvisited node, try to traverse. If a reachable neighbor is already visited, we take that result and perform | for all reachable neighbor
- On each visited node, keep track if it can reach top/left and bottom/right.
- Use a Heap to pick where to start.
**/

var directions = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func pacificAtlantic(heights [][]int) [][]int {
	var res [][]int
	M, N := len(heights), len(heights[0])

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			visited := makeVisited(M, N)
			queue := [][]int{{i, j}}
			tl, br := false, false
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				if visited[cur[0]][cur[1]] {
					continue
				}
				visited[cur[0]][cur[1]] = true
				if cur[0] == 0 || cur[1] == 0 {
					tl = true
				}
				if cur[0] == M-1 || cur[1] == N-1 {
					br = true
				}

				for _, d := range directions {
					x, y := cur[0]+d[0], cur[1]+d[1]
					if x >= 0 && x <= M-1 && y >= 0 && y <= N-1 {
						if heights[cur[0]][cur[1]] >= heights[x][y] && !visited[x][y] {
							queue = append(queue, []int{x, y})
						}
					}
				}
			}

			if tl && br {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func makeVisited(m, n int) [][]bool {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return visited
}
