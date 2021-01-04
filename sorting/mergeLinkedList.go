package sorting

/*
Question 21: https://leetcode.com/problems/merge-two-sorted-lists/

Idea:
 - Noob: Creates an array, store all elements there. and then manually links them again
 - Less noob: do it recursively and let implicit stack do the work : - )

**/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil && l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else if l1 != nil && l2 != nil {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else if l1 != nil {
		return l1
	} else if l2 != nil {
		return l2
	}
	return nil
}
