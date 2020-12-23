package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * Idea:
 * Implement BFS, keeps track of the level
 * by storing len variable independent from
 * the loop.
 */

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var res [][]int
	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		len := len(queue)
		var currentLevel []int
		for i := 0; i < len; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Rights)
			}
			currentLevel = append(currentLevel, queue[0].Val)
			queue = queue[1:]
		}
		res = append(res, currentLevel)
	}
	return res
}
