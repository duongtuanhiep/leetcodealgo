package array

import (
	"container/heap"
)

/*
Question 295: https://leetcode.com/problems/find-median-from-data-stream/

Solution: https://www.geeksforgeeks.org/median-of-stream-of-integers-running-integers/

Idea: if we can maintain 2 equal-sized heap then median will be either avg of 2 root or root of heap with more element.
**/

// An IntHeap is a max-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	Med float64
	Min *MinHeap // Keep all elem greater than med
	Max *MaxHeap // keep all elem less than med
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := &MinHeap{}
	max := &MaxHeap{}
	heap.Init(min)
	heap.Init(max)
	return MedianFinder{0.0, min, max}
}

func (this *MedianFinder) AddNum(num int) {
	if float64(num) > this.Med {
		heap.Push(this.Min, num)
	} else {
		heap.Push(this.Max, num)
	}

	for len(*this.Max)-len(*this.Min) < 0 || len(*this.Max)-len(*this.Min) > 1 {
		if len(*this.Max) > len(*this.Min) {
			heap.Push(this.Min, heap.Pop(this.Max))
		} else {
			heap.Push(this.Max, heap.Pop(this.Min))
		}
	}

	// Calculate new Med
	if len(*this.Max) > len(*this.Min) {
		this.Med = float64((*this.Max)[0])
	} else if len(*this.Max) > 0 && len(*this.Min) > 0 {
		this.Med = (float64((*this.Max)[0]) + float64((*this.Min)[0])) / 2
	} else {
		this.Med = float64(num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	return this.Med
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
