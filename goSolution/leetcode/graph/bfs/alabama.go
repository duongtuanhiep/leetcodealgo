package bfs

/*
Question 993: https://leetcode.com/problems/cousins-in-binary-tree/
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {

	queue := []*TreeNode{root}
	var found1, found2 bool
	for len(queue) > 0 {
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			if queue[0].Val == x {
				found1 = true
			} else if queue[0].Val == y {
				found2 = true
			}

			if queue[0].Left != nil && queue[0].Right != nil {
				if queue[0].Left.Val == x && queue[0].Right.Val == y || queue[0].Left.Val == y && queue[0].Right.Val == x {
					return false
				}
			}

			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		if found1 != found2 {
			break
		}
	}

	return found1 == found2 && found1 == true
}
