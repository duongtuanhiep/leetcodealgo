package realgraphshit

/* Idea:
 * For each available edge available, we can reach the tail
 * from the head, meaning that both head and tail are reachable
 * from the head.
 * All the other element that does not have any edges point to
 * are the element that need to be included since those are
 * unreachable from any other vertices except itself (vertex
 * indegree = 0). These will be the minimum elements.
**/

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	need := make([]int, n)
	var res []int
	for i := range edges {
		need[edges[i][1]]++
	}
	for vertex := 0; vertex < n; vertex++ {
		if need[vertex] == 0 {
			res = append(res, vertex)
		}
	}
	return res
}
