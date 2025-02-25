package graph

/*

Question 2467: https://leetcode.com/problems/most-profitable-path-in-a-tree/

Observation:
- Since len(edges) == len(nodes) - 1 it guaranteed that there will be no cycle
- Since A always travel toward the optimal leaf node, it's guaranteed that there's no backtracking for alice (verifiable with a test case)

Solution:
- Get bob distance and visited node via DFS, make sure to NOT track the node doesn't lead to 0
- Perform BFS on Alice path, calculate income of each step and store it. Once reached a leaf node, compare with Max Income.

**/

func mostProfitablePath(edges [][]int, bob int, amount []int) int {

	// Building travelable map
	pathMap := make(map[int][]int)
	for _, edge := range edges {
		pathMap[edge[0]] = append(pathMap[edge[0]], edge[1])
		pathMap[edge[1]] = append(pathMap[edge[1]], edge[0])
	}

	// Traverse from 0 to Bob to find his path
	bobDist := make(map[int]int)
	bobVisited := make(map[int]struct{})
	bobDfs(pathMap, bob, bobVisited, bobDist)

	queue := []int{0}
	visited := make(map[int]struct{})
	currentIncome := make(map[int]int)
	maxIncome := -100000
	var dist int
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			cur := queue[0]
			queue = queue[1:]
			visited[cur] = struct{}{}

			if _, bobWent := bobVisited[cur]; !bobWent || bobDist[cur] > dist {
				currentIncome[cur] += amount[cur]
			} else if bobDist[cur] == dist {
				currentIncome[cur] += amount[cur] / 2
			}

			hasNext := false
			for _, edge := range pathMap[cur] {
				if _, ok := visited[edge]; ok {
					continue
				}

				queue = append(queue, edge)
				currentIncome[edge] += currentIncome[cur]
				hasNext = true
			}
			if !hasNext && currentIncome[cur] > maxIncome {
				maxIncome = currentIncome[cur]
			}

			n--
		}
		dist++
	}

	return maxIncome
}

// DFS, remove node and ended up false
func bobDfs(pathMap map[int][]int, bob int, visited map[int]struct{}, dist map[int]int) bool {
	visited[bob] = struct{}{}

	if bob == 0 {
		return true
	}

	for _, edge := range pathMap[bob] {
		if _, ok := visited[edge]; ok {
			continue
		}

		dist[edge] = dist[bob] + 1
		if bobDfs(pathMap, edge, visited, dist) {
			return true
		} else {
			delete(visited, edge)
			delete(dist, edge)
		}
	}

	return false
}
