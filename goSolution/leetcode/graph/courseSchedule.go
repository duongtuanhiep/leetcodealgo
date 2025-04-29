package graph

import "slices"

/*

Question 210: https://leetcode.com/problems/course-schedule-ii

Idea: Implement topological sort + cycle detection. This is done by either BFS (Kahn algo) / DFS
Some key consideration for DFS implementation:
- Utilizing the nature of DFS, we can add node to list only after visiting all edges from node. Reversing this becomes topological sort [0,1,2,3]
- When starting a new cycle, instead of appending it to the back of topological sort, we append to front. [4,0,1,2,3]
- Extending visited from 0|1 to 0|1|2 where: 1 is currently visited in cycle and 2 is visited and explored all childs. When encounter a "1" we know theres cycle

Visualitaion:
   0 -> 1 -> 2 -> 3
        ^
		4
**/

func findOrder(numCourses int, prerequisites [][]int) []int {
	topo := make([]int, 0, numCourses)
	edge := make(map[int][]int)

	for _, e := range prerequisites {
		edge[e[1]] = append(edge[e[1]], e[0])
	}

	visited := make([]int, numCourses)
	for i := range numCourses {
		// append to front instead of back
		order, valid := topologicalDfs(visited, edge, i)
		if !valid {
			return []int{}
		}
		slices.Reverse(order)
		topo = append(order, topo...)
	}

	return topo
}

func topologicalDfs(visited []int, edge map[int][]int, cur int) ([]int, bool) {
	if visited[cur] == 2 {
		return nil, true
	}
	if visited[cur] == 1 {
		return nil, false
	}
	visited[cur] = 1

	var order []int
	for _, e := range edge[cur] {
		childOrder, valid := topologicalDfs(visited, edge, e)
		if !valid {
			return nil, false
		}
		order = append(order, childOrder...)
	}

	order = append(order, cur)

	visited[cur] = 2

	return order, true
}
