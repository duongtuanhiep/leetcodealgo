package greedy

import "container/heap"

/*
Question 1354: https://leetcode.com/problems/construct-target-array-with-multiple-sums/

Observation:
- When a number is set we cant change it because future operation will only have higher values.
That imples us when there is a target number available we have to do it. That also implies that
we have to do it sequentially.

- After an operation, we have to "stick" to that number.
**/

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

func isPossible(target []int) bool {
	if len(target) == 1 && target[0] != 1 {
		return false
	}
	minSum := len(target)
	curSum := 0
	h := &IntHeap{}
	heap.Init(h)

	for _, val := range target {
		heap.Push(h, val)
		curSum += val
	}

	for curSum > minSum {
		var last, op, remainder, next int
		last = heap.Pop(h).(int)
		op, remainder = last/(curSum-last), last%(curSum-last)
		if remainder == 0 && op > 1 {
			next = last - (curSum-last)*(op-1)
		} else if remainder > 0 && op == 0 {
			return false
		} else {
			next = remainder
		}
		if next < 1 {
			return false
		}
		curSum += next - last
		heap.Push(h, next)
	}
	return curSum == minSum
}

/*
test case
[1,1,1,2]
[2,999]
[1,1000000000]
[9,9,9]
**/
