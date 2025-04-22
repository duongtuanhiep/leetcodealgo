package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	nodes := treeDfs(root, nodes)
	// reassemble

	for i := range nodes {
		nodes[i].Left = nil
		if i != len(nodes)-1 {
			nodes[i].Right = nodes[i+1]
		} else {
			nodes[i].Right = nil
		}
	}
	return nodes[0]
}

func treeDfs(node *TreeNode) []*TreeNode {
	var res []*TreeNode

	if node != nil {
		if node.Left != nil {
			res = append(res, treeDfs(node.Left)...)
		}
		res = append(res, node)
		if node.Right != nil {
			res = append(res, treeDfs(node.Right)...)
		}
	}
	return res
}
