package bfs

/*
Question 637: https://leetcode.com/problems/average-of-levels-in-binary-tree/

Idea: BFS, keep track of the level

**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		nodeNum := len(queue)
		var total int
		for i := 0; i < nodeNum; i++ {
			total += queue[0].Val
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res, float64(total)/float64(nodeNum))
	}

	return res
}
