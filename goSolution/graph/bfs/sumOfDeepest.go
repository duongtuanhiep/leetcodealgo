package bfs

/*
Question 1302: https://leetcode.com/problems/deepest-leaves-sum/


**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	levelSum := 0
	for curLen := len(queue); curLen != 0; curLen = len(queue) {
		levelSum = 0
		for i := 0; i < curLen; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			levelSum += queue[0].Val
			queue = queue[1:]
		}
	}
	return levelSum
}
