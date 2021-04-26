package greedy

import (
	"container/heap"
)

/*
Question 1642: Furthest can reach
https://leetcode.com/problems/furthest-building-you-can-reach/

Test case example:
[4,3,11,3,5,6,3,2,4,3,7,6,9,14,12]
7
2
--
[14,3,19,3]
17
0
--
[4,12,2,7,3,18,20,3,19]
10
2

Idea: Greedy-ish
We try to use brick as much as we can, if it is not possible replace the most expensive
brick operation with a ladder and then continue.
Create a heap to maintain the most expensive operation so far.
*/

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &IntHeap{}
	heap.Init(h)
	var res int
	for i := 1; i < len(heights); i++ {
		if cost := diff(heights[i-1], heights[i]); cost > 0 {
			heap.Push(h, cost)
			bricks -= cost
		}
		if bricks < 0 {
			if ladders > 0 {
				bricks += heap.Pop(h).(int)
				ladders--
			} else {
				break
			}
		}
		res = i
	}
	return res
}

func diff(a, b int) int {
	if a > b {
		return 0
	}
	return b - a
}
