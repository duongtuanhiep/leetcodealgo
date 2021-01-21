package greedy

import (
	"container/heap"
)

/*
Question 1673: https://leetcode.com/problems/find-the-most-competitive-subsequence/

// Algorithm (greedy-ish)
From Lastadded -> len(nums) - k - len(cur): Get Max element of array

my solution explained:
- Build a heaps to store min, a hash map to store position of min element
- On each iteration, checks if min with index i satify:
	1) if i < last: pop from heap, removes from map
	2) if i > len - k (elem remain in result arr): pop, add to queue to later added back to arr
	3) if last < i <= len - k: pop removes from map, add to arr
- Add everything from 2) back to heap
Repeat until k == 0

Optimization:
- appending everyhing left lastAddedPos = len - k - 1
**/
func mostCompetitive(nums []int, k int) []int {

	if len(nums) == k {
		return nums
	}

	var res []int
	var heapNums IntHeap
	heapNums = make([]int, len(nums))
	copy(heapNums, nums)
	heap.Init(&heapNums)

	pos := make(map[int][]int)
	for i, num := range nums {
		if _, found := pos[num]; !found {
			pos[num] = []int{i}
		} else {
			pos[num] = append(pos[num], i)
		}
	}

	lastPos := -1
	queue := []int{}
	for k > 0 {
		var minNum int
		for {
			minNum = heapNums[0]          // get min
			heap.Pop(&heapNums)           // pop from heap
			if pos[minNum][0] > lastPos { // check leftbound
				if pos[minNum][0] <= len(nums)-k { // checks right bound
					lastPos, pos[minNum] = pos[minNum][0], pos[minNum][1:]
					res = append(res, minNum) // append Arr
					break
				} else {
					queue = append(queue, minNum) // add to queue
				}
			} else {
				pos[minNum] = pos[minNum][1:] // drop from arr
			}
		}
		k-- // reduce counter

		// pushing back
		for len(queue) > 0 {
			heap.Push(&heapNums, queue[0])
			queue = queue[1:]
		}

		// Minor optimization
		if lastPos == len(nums)-k-1 {
			res = append(res, nums[lastPos+1:]...)
			break
		}
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
