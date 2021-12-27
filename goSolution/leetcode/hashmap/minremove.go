package hashmap

import "container/heap"

/*
Question 1338:https://leetcode.com/problems/reduce-array-size-to-the-half/

**/

func minSetSize(arr []int) int {
	counter := make(map[int]int)
	cur := len(arr)
	for _, val := range arr {
		counter[val]++
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, num := range counter {
		heap.Push(h, num)
	}

	last := len(*h)
	for cur > len(arr)/2 {
		cur -= heap.Pop(h).(int)
	}

	return last - len(*h)
}

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
