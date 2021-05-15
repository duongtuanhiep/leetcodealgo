package linkedlist

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	// base case
	holder := root.Right
	if root.Left != nil {
		root.Right = root.Left
	}
	if root.Right != nil {

	}
}
