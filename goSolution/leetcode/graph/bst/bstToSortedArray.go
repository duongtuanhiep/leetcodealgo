package bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Observation: Since it is a BST tree, meaning left < right.
 * A single DFS pass Inorder will be albe to solve this
 *
 *
 * Since this is a BST, every node if it has left child or right
 * child will have values follow the order :
 * node.Left.Val < node.Val < node.Right.Val
 * This suggest an implementation where you can visit every node
 * and add follows order: left -> mid -> right (inorder)
 *
 * Perharps there is a way to piggy back on the previous traversal
 * without having an array to store, putting the node childs while
 * traversing in order ?
**/

// Implement using recursive approach with an arrays
func increasingBST(root *TreeNode) *TreeNode {

	// implement preorder DFS
	var stacks []*TreeNode
	sortedStack := traverseInDFS(root, stacks)
	for pos := range sortedStack {
		sortedStack[pos].Left = nil
		if pos == len(sortedStack)-1 {
			sortedStack[pos].Right = nil
			break
		}
		sortedStack[pos].Right = sortedStack[pos+1]
	}
	return sortedStack[0]
}

// func traverseInDFS(node *TreeNode, stacks []*TreeNode) []*TreeNode {
// 	if node.Left != nil {
// 		stacks = traverseInDFS(node.Left, stacks)
// 	}
// 	stacks = append(stacks, node)
// 	if node.Right != nil {
// 		stacks = traverseInDFS(node.Right, stacks)
// 	}
// 	return stacks
// }

func smarterDFS(node *TreeNode, holder *TreeNode) *TreeNode { // Try to return the "parent" node with smaller num
	if node.Left != nil {
		holder = node.Left
		holder.Right = node
		smarterDFS(node.Left, holder)
	}
	stacks = append(stacks, node)
	if node.Right != nil {
		stacks = traverseInDFS(node.Right, stacks)
	}
}
