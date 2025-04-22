package heap

import (
	"container/heap"
	"sort"
)

/*
Greedy: for each day d, try to do the event that end the earlest if you havent

**/

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := &MinHeap{}
	heap.Init(pq)

	i, day, count := 0, 1, 0
	for i < len(events) || pq.Len() > 0 {
		// Add all events that start at or before "day"
		for i < len(events) && events[i][0] <= day {
			heap.Push(pq, events[i])
			i++
		}

		// Remove events that have already ended
		for pq.Len() > 0 && (*pq)[0][1] < day {
			heap.Pop(pq)
		}

		// Attend the event that ends the earliest
		if pq.Len() > 0 {
			heap.Pop(pq)
			count++
			day++
		} else {
			// If there are no events to attend today, move to the start day of the next event
			if i < len(events) {
				day = events[i][0]
			} else {
				break // No more events
			}
		}

	}

	return count
}

// An IntHeap is a min-heap of ints.
type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i][1], h[j][1] = h[j][1], h[i][1] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
