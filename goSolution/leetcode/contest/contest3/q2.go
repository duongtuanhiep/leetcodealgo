package contest3

import "container/heap"

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

type SeatManager struct {
	H *IntHeap
}

func Constructor(n int) SeatManager {
	h := &IntHeap{}
	heap.Init(h)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.H).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.H, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
