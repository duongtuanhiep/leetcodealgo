package graph

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 *
 * Idea: Do a "mirror" BFS on left half and right half, adds all the nodes
 * into left arr and right arr. Update nil nodes with nil value in array
 * and then compare those elements
 *
 */

/** Representing a tree with an array
 * You've seen two approaches to implementing a sequence data structure:
 * either using an array, or using linked nodes. We extended our idea of
 * linked nodes to implement a tree data structure. It turns out we can
 * also use an array to represent a tree.
 * Here's how we implement a binary tree:
 * The root of the tree will be in position 1 of the array (nothing is at
 * position 0). We can define the position of every other node in the tree
 * recursively:
 * The left child of a node at position n is at position 2n.
 * The right child of a node at position n is at position 2n + 1.
 * The parent of a node at position n is at position n/2.
**/

func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	var left []*TreeNode
	var right []*TreeNode

	if root == nil {
		return true
	} else if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	} else if root.Left == root.Right && root.Left == nil {
		return true
	}

	// Traverse left
	left = append(left, root.Left)
	queue = append(queue, root.Left)
	for len(queue) > 0 {
		cur := queue[0]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			left = append(left, cur.Left)
		} else {
			left = append(left, nil)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			left = append(left, cur.Right)
		} else {
			left = append(left, nil)
		}
		queue = queue[1:]
	}

	// Traverse right
	right = append(right, root.Right)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		cur := queue[0]
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			right = append(right, cur.Right)
		} else {
			right = append(right, nil)
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			right = append(right, cur.Left)
		} else {
			right = append(right, nil)
		}
		queue = queue[1:]
	}

	for pos := range left {
		if left[pos] == nil && right[pos] == nil {
			return true
		}
		if right[pos] == nil || left[pos] == nil {
			return false
		}
		if left[pos].Val != right[pos].Val {
			return false
		}
	}

	return true
}
