package bfs

/*
Question 450: https://leetcode.com/problems/delete-node-in-a-bst/
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
Idea: since its unique just simply perform binary search and moves nodes accordingly.
Bubbling the needed nodes to the bottom and then remove it.
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
	dummyHead := &TreeNode{-1, root, nil}

	for root != nil {
		if root.Val == key {
			root.Left = 
		}

	}

	return dummyHead.Left
}
