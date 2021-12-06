package linkedlist

/*
Question 430: https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list


Idea:
- Easy: Create an array to store node in order, create a stack to mantain relative
order of Next and Child List node. After we have the list node we can manually link them back
together

- Not so easy: Instead of creating array to maintain final order, we can try to create
a dump node and handle the linking in places. The local order has already been maintained
by the stack

- Alternative: instead of using stack to maintain the order we can delegate that to
the program stack by using recusrion call.

**/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	var cur, prev *Node
	stack, dump := []*Node{root}, &Node{0, nil, nil, root}
	for len(stack) != 0 {
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if cur.Next != nil {
			stack = append(stack, cur.Next)
		}
		if cur.Child != nil {
			stack = append(stack, cur.Child)
		}

		if len(stack) != 0 {
			cur.Next = stack[len(stack)-1]
		} else {
			cur.Next = nil
		}
		cur.Child, cur.Prev, prev = nil, prev, cur
	}

	return dump.Child
}
