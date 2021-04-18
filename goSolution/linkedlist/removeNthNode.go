package linkedList

import "fmt"

/*
Question 19: https://leetcode.com/problems/remove-nth-node-from-end-of-list/


Idea:
Easy solution:
- Creates an array to store reference, then just manually moves the n+1 from end Next to other.



**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	if n+1 <= len(nodes) && n-1 > 0 {
		nodes[len(nodes)-n-1].Next = nodes[len(nodes)-n+1]
	} else if n == 1 {
		nodes[len(nodes)-2].Next = nil
	} else if n == len(nodes) {
		return nodes[1]
	}

	return nodes[0]
}
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	holderNode := &ListNode{head.Val, head.Next}
	dummyHead := ListNode{-1, holderNode}
	for head != nil || holderNode != nil {
		// fmt.Println("RIP")
		if n > 0 {
			n--
			head = head.Next
		} else if n == 0 && head == nil {
			fmt.Println(holderNode)
			n--
			// Do swapping here
			if holderNode.Next != nil {
				holderNode.Val = holderNode.Next.Val
				holderNode.Next = holderNode.Next.Next
			} else if holderNode.Next == nil {
				holderNode = nil
			}
			fmt.Println(holderNode)
		} else {
			if head != nil {
				head = head.Next
			}
			holderNode = holderNode.Next
		}
	}
	return dummyHead.Next
}
