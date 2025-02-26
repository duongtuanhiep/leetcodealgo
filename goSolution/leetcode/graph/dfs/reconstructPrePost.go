package dfs

/*

Question 889: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

Couple observation:
- binary tree: left < right

**/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	node := &TreeNode{Val: preorder[0]}

	posPre := make(map[int]int)
	for i, val := range preorder {
		posPre[val] = i
	}

	posPost := make(map[int]int)
	for i, val := range postorder {
		posPost[val] = i
	}
	// fmt.Println(preorder, postorder)
	// Findind possible left subtree and right subtree
	left, right := preorder[1], postorder[len(postorder)-2]
	if left != right { // current node has both left and right
		node.Left, node.Right = constructFromPrePost(preorder[1:posPre[right]], postorder[:posPost[left]+1]), constructFromPrePost(preorder[posPre[right]:], postorder[posPre[left]:len(postorder)-1])
	} else {
		node.Left = constructFromPrePost(preorder[1:], postorder[:len(postorder)-1])
	}

	return node
}
