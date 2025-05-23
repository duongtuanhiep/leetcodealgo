package slidingwindows

import "container/heap"

/*
Question 239: https://leetcode.com/problems/sliding-window-maximum/

*/

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

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	h := &IntHeap{}
	heap.Init(h)
	lastPos := make(map[int]int)
	for i := 0; i < k; i++ {
		lastPos[nums[i]] = i
		heap.Push(h, nums[i])
	}
	res = append(res, (*h)[0])

	for i := k; i < len(nums); i++ {
		lastPos[nums[i]] = i
		heap.Push(h, nums[i])
		for lastPos[(*h)[0]] < i-k+1 {
			heap.Pop(h)
		}
		res = append(res, (*h)[0])
	}
	return res
}
