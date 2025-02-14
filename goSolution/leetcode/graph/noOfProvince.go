package graph

/*
Question 547: https://leetcode.com/problems/number-of-provinces/

**/

func findCircleNum(isConnected [][]int) int {
	var res int
	visited := make([]bool, len(isConnected))
	for i := range isConnected {
		if !visited[i] {
			var queue []int = []int{i}
			for len(queue) > 0 {
				cur := queue[0]
				for j := range isConnected[cur] {
					if isConnected[cur][j] == 1 && !visited[j] {
						queue = append(queue, j)
					}
				}
				visited[cur] = true
				queue = queue[1:]
			}
			res++
		}
	}
	return res
}
