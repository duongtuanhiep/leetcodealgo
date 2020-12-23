package bfs

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Idea: Implement BFS, use the current length of the queue to maintain the level.
// put a length variable to main for loop. only interacts with node originally belongs here

func connect(root *Node) *Node {

	// base case: check root
	if root == nil {
		return root
	}

	// Init queue
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if i == l-1 {
				queue[0].Next = nil
			} else {
				queue[0].Next = queue[1]
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}

	return root
}
