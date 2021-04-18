package linkedlist

/*
Question 86: https://leetcode.com/problems/partition-list/

Idea:
- Rough: create 2 list of elements left and right of pivots. Manually "link" them together
	(last element of left to first of right)
- Smarter: Create 2 fake notes, to hold leftmost and rightMost elements. Link them together
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	leftDummy, rightDummy := &ListNode{-101, nil}, &ListNode{-101, nil}
	newRoot, newRootRight := ListNode{-101, leftDummy}, ListNode{-101, rightDummy}
	for head != nil {
		if head.Val < x {
			leftDummy.Next = head
			leftDummy = leftDummy.Next
		} else {
			rightDummy.Next = head
			rightDummy = rightDummy.Next
		}
		head = head.Next
	}

	// Reattach middle element and last element
	if leftDummy.Val > -100 && newRootRight.Next.Val > -100 {
		leftDummy.Next = newRootRight.Next
	}
	if rightDummy.Val > -100 {
		rightDummy.Next = nil
	}

	if leftDummy.Val > -100 {
		return newRoot.Next
	}
	return newRootRight.Next
}
