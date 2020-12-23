package graph

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traverse(0, root)
}

func traverse(depth int, node *TreeNode) int {
	if node == nil {
		return depth
	}
	depth++
	depthLeft := traverse(depth, node.Left)
	depthRight := traverse(depth, node.Right)
	if depthLeft > depthRight {
		return depthLeft
	}
	return depthRight
}
