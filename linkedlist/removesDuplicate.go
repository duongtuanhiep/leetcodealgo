package linkedlist

/**
Question 82: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

* Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

Idea:
- Create an holder arrays and a fucking map. return first element of map
-
*/
// func deleteDuplicates(head *ListNode) *ListNode {
// 	arrNode, reArrNode, nodeMap := []*ListNode{}, []*ListNode{}, make(map[int]int)
// 	if head == nil {
// 		return nil
// 	} else {
// 		for head != nil {
// 			arrNode = append(arrNode, head)
// 			nodeMap[head.Val]++
// 			head = head.Next
// 		}
// 	}

// 	for i := range arrNode {
// 		if nodeMap[arrNode[i].Val] == 1 {
// 			reArrNode = append(reArrNode, arrNode[i])
// 		}
// 	}

// 	for i := range reArrNode {
// 		if i != len(reArrNode)-1 {
// 			reArrNode[i].Next = reArrNode[i+1]
// 			continue
// 		}
// 		reArrNode[i].Next = nil
// 	}

// 	if len(reArrNode) > 0 {
// 		return reArrNode[0]
// 	}
// 	return nil
// }

// Slightly optimized
func deleteDuplicates(head *ListNode) *ListNode {
	var holder []*ListNode
	var lastVal int = 101
	for head != nil {
		if lastVal == head.Val {
			if len(holder) > 0 && holder[len(holder)-1].Val == lastVal {
				holder = holder[:len(holder)-1]
			}
		} else {
			holder = append(holder, head)
			lastVal = head.Val
		}
		head = head.Next
	}

	for i := range holder {
		if i == len(holder)-1 {
			holder[i].Next = nil
			break
		}
		holder[i].Next = holder[i+1]
	}

	if len(holder) > 0 {
		return holder[0]
	}
	return nil
}
