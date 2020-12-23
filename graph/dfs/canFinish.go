package dfs

// Node ddmm
type Node struct {
	Value    int
	ParentTo *Node
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	visited := map[int]int{}

	// Preprocessing
	courseGraph := make([][]int, numCourses)

	for _, pair := range prerequisites {
		if len(courseGraph[pair[0]]) == 0 {
			courseGraph[pair[0]] = []int{pair[1]}
		}
		courseGraph[pair[0]] = append(courseGraph[pair[0]], pair[1])

	}

	var result = true
	for i := 0; i < len(courseGraph); i++ {
		result = traverse(i, courseGraph, visited) && result
	}

	return result
}

func traverse(node int, nodeArr [][]int, visited map[int]int) bool {
	visited[node] = 1
	for j := 0; j < len(nodeArr[node]); j++ {
		if visited[nodeArr[node][j]] == 1 {
			return false
		} else if visited[nodeArr[node][j]] == 0 {
			return traverse(nodeArr[node][j], nodeArr, visited)
		}
	}
	visited[node] = 2
	return true
}
