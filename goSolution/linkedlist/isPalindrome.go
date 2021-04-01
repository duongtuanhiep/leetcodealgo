package linkedlist

/**
Question 234: https://leetcode.com/problems/palindrome-linked-list/



Definition for singly-linked list.
type ListNode struct {
      Val int
      Next *ListNode
}

Solution:
- One easy solution is to save the whole list to a place holder. THen just check from middle or 2 ends and compare
- Try to reverse the first half somehow, and then compare it with the second half.
*/

/*
func isPalindrome(head *ListNode) bool {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	var i, j int
	if len(list)%2 == 0 {
		i, j = len(list)/2-1, len(list)/2
	} else {
		i, j = len(list)/2-1, len(list)/2+1
	}
	for i >= 0 {
		if list[i] != list[j] {
			return false
		}
		i, j = i-1, j+1
	}
	return true
}
**/

func isPalindrome(head *ListNode) bool {
	// base
	if head.Next == nil {
		return true
	}

	var prev, slow, fast, next *ListNode
	prev, slow, fast, next = nil, head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next = next.Next
		slow.Next = prev
		prev = slow
		slow = next

	}

	if fast != nil {
		next = next.Next
	}

	for prev != nil && next != nil {
		if prev.Val != next.Val {
			return false
		}
		prev, next = prev.Next, next.Next
	}
	return true
}
