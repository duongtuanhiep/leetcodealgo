package sorting

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	// Copy of original
	cpyListNode := head
	// Preprocessing
	var listNode []*ListNode
	var listVal []int

	if head.Next == nil {
		return head
	}

	for head != nil {
		listNode = append(listNode, cpyListNode)
		listVal = append(listVal, head.Val)
		head = head.Next
	}

	// insertion sort
	for pos, val := range listNode {
		if pos == 0 {
			continue
		}
		value := val
		hole := pos

		for hole > 0 && listNode[hole-1].Val > value.Val {
			listNode[hole] = listNode[hole-1]
			hole--
		}
		listNode[hole] = value
	}

	for pos := range listNode {
		fmt.Println(listNode[pos].Val)
	}

	for pos := range listNode {
		if pos+1 == len(listNode) {
			listNode[pos].Next = nil
			break
		}
		listNode[pos].Next = listNode[pos+1]
	}

	return listNode[0]
}
