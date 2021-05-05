package bfs

func findBottomLeftValue(root *TreeNode) int {
	var queue []*TreeNode
	var res int
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		res := queue[0].Val
		for i := 0; i < length; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return res
}
