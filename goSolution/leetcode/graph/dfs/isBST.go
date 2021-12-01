package dfs

/*
Question 98: https://leetcode.com/problems/validate-binary-search-tree/
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var nodes []int
	nodes = traverseAsc(root)
	last := nodes[0]
	for _, val := range nodes[1:] {
		if last >= val {
			return false
		} else {
			last = val
		}
	}
	return true
}

func traverseAsc(root *TreeNode) []int {
	var res []int
	if root.Left != nil {
		res = traverseAsc(root.Left)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		rightSide := traverseAsc(root.Right)
		res = append(res, rightSide...)
	}
	return res
}
