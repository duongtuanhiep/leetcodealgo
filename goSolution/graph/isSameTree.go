package graph

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 * Idea:
 * 1. Run through p and q via bfs, creates an arr to
 * store both nodes. Run through arr again and compares
 *
 * 2. Recursively go through both nodes via DFS. Any
 * traversal order will works.
 */

// DFS Solution
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	if left == false || right == false {
		return false
	}
	return true
}
