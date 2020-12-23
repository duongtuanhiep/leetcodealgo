package dfs

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/** Max Path Idea:
Implement DFS and keep tracks of the value. Created a "local" max
and a "global" max represents all the path.
*/

func maxPathSum(root *TreeNode) int {

	// return int(math.Max(float64(root.Val+maxLeft+maxRight), math.Max(float64(maxLeft), float64(maxRight))))
	return maxPathDFS(root, 0)
}

func maxPathDFS(node *TreeNode, maxTillNow int) int {

	// If all the maxPath until this point is negative, better to start new
	if maxTillNow <= 0 {
		maxTillNow = node.Val
	} else {
		maxTillNow += node.Val
	}

	// DFS left, DFS right
	var maxLeft, maxRight int
	if node.Left != nil && node.Right != nil {
		maxLeft = maxPathDFS(node.Left, maxTillNow)
		maxRight = maxPathDFS(node.Right, maxTillNow)
		return int(math.Max(float64(maxLeft+maxRight-maxTillNow), math.Max(float64(maxLeft), float64(maxRight))))
		// return int(math.Max(float64(maxLeft), float64(maxRight)))
	} else if node.Right != nil {
		maxRight = maxPathDFS(node.Right, maxTillNow)
		// return int(math.Max(float64(maxTillNow+maxRight), float64(maxRight)))
		return maxRight
	} else if node.Left != nil {
		maxLeft = maxPathDFS(node.Left, maxTillNow)
		// return int(math.Max(float64(maxTillNow+maxLeft), float64(maxLeft)))
		return maxLeft

	}
	return maxTillNow
}
