package bfs

import (
	"container/heap"
)

/*
Question 2503: https://leetcode.com/problems/maximum-number-of-points-from-grid-queries

Idea: Instead of perform BFS for each queries, we can possibly try to perform only once BFS and save the states somehow ?

We know the max values of a single cell (1 <= x <= 1e6), we can perform:
- For each threshold value (1 -> 1e6), we try to perform a BFS and record the score up until we can do it no longer.
- Once we have exhausted all the possible option for a certain options, we then adjust threshold to the next higher score.
- We repeat first step, by the end we would have traversed the grid and have an array that store possible points for each threshold
- We merge that with the queries costs.

**/

func maxPoints(grid [][]int, queries []int) []int {
	maxVisited := maxNum(queries)
	thresholdPoint := make([]int, 1e6+1)
	var visited [][]bool
	for i := range grid {
		visited = append(visited, make([]bool, len(grid[i])))
	}

	minHeap := &MinHeap{node{x: 0, y: 0, val: grid[0][0]}}
	heap.Init(minHeap)
	pts := 0
	for i := 1; i <= 1e6; i++ { // Optimize this by only trying the queries values instead of a natural range
		for tmp := *minHeap; len(*minHeap) > 0 && i > tmp[0].val; tmp = *minHeap {
			cur := heap.Pop(minHeap).(node)
			if visited[cur.x][cur.y] {
				continue
			}
			if cur.x > 0 && !visited[cur.x-1][cur.y] {
				heap.Push(minHeap, node{x: cur.x - 1, y: cur.y, val: grid[cur.x-1][cur.y]})
			}
			if cur.y > 0 && !visited[cur.x][cur.y-1] {
				heap.Push(minHeap, node{x: cur.x, y: cur.y - 1, val: grid[cur.x][cur.y-1]})
			}
			if cur.x < len(grid)-1 && !visited[cur.x+1][cur.y] {
				heap.Push(minHeap, node{x: cur.x + 1, y: cur.y, val: grid[cur.x+1][cur.y]})
			}
			if cur.y < len(grid[0])-1 && !visited[cur.x][cur.y+1] {
				heap.Push(minHeap, node{x: cur.x, y: cur.y + 1, val: grid[cur.x][cur.y+1]})
			}

			visited[cur.x][cur.y] = true
			pts++
		}

		thresholdPoint[i] = pts
		if i == maxVisited {
			break
		}
	}

	var res []int
	for _, val := range queries {
		res = append(res, thresholdPoint[val])
	}

	return res
}

func maxNum(nums []int) int {
	res := 0
	for i := range nums {
		if nums[res] < nums[i] {
			res = i
		}
	}
	return nums[res]
}

type node struct {
	x, y int
	val  int
}

// An IntHeap is a min-heap of ints.
type MinHeap []node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
