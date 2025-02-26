package bfs

/*
Question 116: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
**/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	temp := &Node{-1, nil, nil, root}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		var last *Node
		for i := 0; i < length; i++ {
			cur := queue[0]
			if last != nil {
				last.Next = cur
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			queue = queue[1:]
			last = cur
		}
	}

	return temp.Next
}
