package dfs

/*

Question 1123: https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/?envType=daily-question&envId=2025-04-04

Essentially tries to return LCA:
- If there's 2, return LCA.
- If there's one, return itself.
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _ := getLCA(root)
	return node
}

func getLCA(cur *TreeNode) (*TreeNode, int) {
	if cur == nil {
		return nil, 0
	}

	lcaLeft, leftDepth := getLCA(cur.Left)
	lcaRight, rightDepth := getLCA(cur.Right)

	if leftDepth > rightDepth {
		return lcaLeft, leftDepth + 1
	}
	if leftDepth < rightDepth {
		return lcaRight, rightDepth + 1
	}

	return cur, leftDepth + 1
}
