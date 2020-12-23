package realgraphshit

/* Idea:
 * Graph branches out and will always go to final elements
 * We can try to create every possible routes from a points.
 * Return the array to function who calls it and appends to with whatever elements.
 * This works very similar to DFS.
**/
func allPathsSourceTarget(graph [][]int) [][]int {
	return *discover(graph, 0)
}

func discover(graph [][]int, node int) *[][]int {
	var routesFromNode [][]int
	if node == len(graph)-1 {
		return &[][]int{{node}}
	}

	for i := range graph[node] {
		subRoute := discover(graph, graph[node][i])
		for _, path := range *subRoute {
			route := []int{node}
			route = append(route, path...)
			routesFromNode = append(routesFromNode, route)
		}
	}

	return &routesFromNode
}
