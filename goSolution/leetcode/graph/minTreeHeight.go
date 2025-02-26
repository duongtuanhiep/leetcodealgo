package graph

/*
Question 310: https://leetcode.com/problems/minimum-height-trees/

**/

func findMinHeightTrees(n int, edges [][]int) []int {
	nodes := make([][]int, n)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}

	var minHeight []int
	curMinHeight := 2001
	// bfs
	for i := range nodes {
		var queue []int
		visited := make(map[int]bool)
		queue = append(queue, i)
		curHeight := 0
		for len(queue) > 0 {
			curHeight++
			length := len(queue)
			for j := 0; j < length; j++ {
				curNodes := queue[0]
				visited[curNodes] = true
				queue = queue[1:]
				for _, next := range nodes[curNodes] {
					if !visited[next] {
						queue = append(queue, next)
					}
				}
			}
		}

		if curHeight < curMinHeight {
			curMinHeight = curHeight
			minHeight = minHeight[:0]
			minHeight = append(minHeight, i)
		} else if curHeight == curMinHeight {
			minHeight = append(minHeight, i)
		}
	}

	return minHeight
}
