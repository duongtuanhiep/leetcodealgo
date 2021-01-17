package dfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func inorderTraversal(root *TreeNode) []int {
// 	var result []int

// 	// add node to int if have to backtrack from left
// 	if root.Left != nil {
// 		result = append(result, inorderTraversal(root.Left, result)...)
// 	} else if root.Right != nil {
// 		result = append(result, root.Val)
// 		result = append(result, inorderTraversal(root.Right, result)...)
// 	}
// 	return result
// }

// func toNodeList
