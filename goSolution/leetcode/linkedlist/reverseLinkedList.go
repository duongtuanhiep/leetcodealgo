package linkedlist

/*
Question 216: https://leetcode.com/problems/reverse-linked-list/

Idea:
Easy: Create a stack.
Medium: Iteratively switch & swap node
Sample: https://www.geeksforgeeks.org/reverse-a-linked-list/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	holder := []*ListNode{}
	for head != nil {
		holder = append(holder, head)
		head = head.Next
	}

	for i := len(holder) - 1; i >= 0; i-- {
		if i != 0 {
			holder[i].Next = holder[i-1]
		} else {
			holder[i].Next = nil
		}
	}

	return holder[len(holder)-1]
}
*/

func reverseList(head *ListNode) *ListNode {
	// res := &ListNode{0, head}
	var next, prev *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
