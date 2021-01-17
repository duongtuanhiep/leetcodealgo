package bitmanipulation

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var bits []int
	total := 0
	// preprocessing N
	for head != nil {
		bits = append(bits, head.Val)
		head = head.Next
	}

	// calculate total
	for pos, _ := range bits {
		total += bits[len(bits)-pos-1] * int(math.Pow(float64(2), float64(pos)))
	}
	return total
}
