package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
// Perform Binary Search on BST
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root.Val >= p.Val && root.Val >= q.Val || root.Val <= p.Val && root.Val <= q.Val {
		if root.Val == p.Val || root.Val == q.Val {
			return root
		}
		if root.Val > p.Val && root.Val > q.Val {
			if root.Right == nil {
				return root
			}
			root = root.Right
		} else {
			if root.Left == nil {
				return root
			}
			root = root.Left
		}
	}

	return root
}
