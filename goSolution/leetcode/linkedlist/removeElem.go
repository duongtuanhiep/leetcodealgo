package linkedlist

/*
Question 203: https://leetcode.com/problems/remove-linked-list-elements/
/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

func removeElements(head *ListNode, val int) *ListNode {
	var last, cur *ListNode = nil, head
	dummyHead := &ListNode{-1, last}
	for cur != nil {
		if cur.Val != val {
			last, last.Next = cur, cur.Next
		}
		last = last.Next
		cur = cur.Next
	}

	return last.Next
}
