package dfs

/*
Question 563: https://leetcode.com/problems/binary-tree-tilt/
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	var tiltLeft, tiltRight int
	var sumLeft, sumRight int
	if root.Left != nil {
		tiltLeft = findTilt(root.Left)
		sumLeft = getSum(root.Left)
	}
	if root.Right != nil {
		tiltRight = findTilt(root.Right)
		sumRight = getSum(root.Right)
	}
	return abs(sumLeft-sumRight) + tiltLeft + tiltRight
}

func getSum(root *TreeNode) int {
	var left, right int
	if root.Left != nil {
		left = getSum(root.Left)
	}
	if root.Right != nil {
		right = getSum(root.Right)
	}
	return root.Val + left + right
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
