package heap

import "container/heap"

/*
Question 3066: https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii

Put everything into min heap, pop 2 elements and do the required computation.
**/

func minOperations(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}

	var res int
	for h.Len() > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		if a >= k {
			return res
		}
		heap.Push(h, a*2+b)
		res++
	}

	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
