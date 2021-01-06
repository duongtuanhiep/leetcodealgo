package linkedlist

import "math/rand"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// type Solution struct {
// 	Nodes []*ListNode
// }

// /** @param head The linked list's head.
//   Note that the head is guaranteed to be not null, so it contains at least one node. */
// func Constructor(head *ListNode) Solution {
// 	var res Solution
// 	for head != nil {
// 		res.Nodes = append(res.Nodes, head)
// 		head = head.Next
// 	}
// 	return res
// }

// /** Returns a random node's value. */
// func (this *Solution) GetRandom() int {
// 	return this.Nodes[rand.Intn(len(this.Nodes))].Val
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

// Algorithm R:  Randomize without keeping track: O(n)
type Solution struct {
	*ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	var res Solution
	res.ListNode = head
	return res
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	output := this.ListNode.Val
	node := this.ListNode
	counter := 1
	for node != nil {
		if rand.Intn(counter) == 0 {
			output = node.Val
		}
		node = node.Next
		counter++
	}
	return output
}
