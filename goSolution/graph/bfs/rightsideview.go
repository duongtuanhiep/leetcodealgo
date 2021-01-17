package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

Idea: piggy back on BFS implementation and keep track of current
layer (hence the i-> previous len)
Be a little subtle and enqueue the right part (the front part) in first
*/

func rightSideView(root *TreeNode) []int {
	var res []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		res = append(res, queue[0].Val)
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[0].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			queue = queue[1:]
		}
	}
	return res
}
