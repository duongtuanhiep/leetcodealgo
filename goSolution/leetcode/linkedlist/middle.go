package linkedlist

/*
Question 876: https://leetcode.com/problems/middle-of-the-linked-list/
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	h1, h2 := head, head
	step := 0
	for h2 != nil {
		if step%2 == 1 {
			h1 = h1.Next
		}
		h2 = h2.Next
		step++
	}
	return h1
}
