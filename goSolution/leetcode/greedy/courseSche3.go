package greedy

import "container/heap"

import "sort"

/*
Question 630: https://leetcode.com/problems/course-schedule-iii/

Idea:
During iteration, say I want to add the current course, currentTotalTime being total time of all courses taken till
now, but adding the current course might exceed my deadline or it doesn’t.
1. If it doesn’t, then I have added one new course. Increment the currentTotalTime with duration of current course.
2. If it exceeds deadline, I can swap current course with current courses that has biggest duration.
- No harm done and I might have just reduced the currentTotalTime, right?
- What preprocessing do I need to do on my course processing order so that this swap is always legal?

Solution:
- Create a max heap to store the longest duration course taken
- Create a "duration" to denote the time taken to finish all the course so far
- Preprocessing: Sort all the course in accending order by its deadline, so that we make sure that we dont "miss"
courses that has nearer deadline because we have added some other course in previously.

Taken from solution online:

From the above example, we can conclude that it is always profitable to take the course with a smaller end day prior
to a course with a larger end day. This is because, the course with a smaller duration, if can be taken, can surely
be taken only if it is taken prior to a course with a larger end day.
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

func scheduleCourse(courses [][]int) int {
	sort.SliceStable(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })

	var duration, taken int
	weight := &IntHeap{0}
	heap.Init(weight)
	for _, course := range courses {
		if duration+course[0] <= course[1] {
			heap.Push(weight, course[0])
			duration += course[0]
			taken++
		} else if (*weight)[0] >= course[0] && duration+course[0]-(*weight)[0] <= course[1] { // Decide if it can be swapped - if yes
			duration += course[0] - heap.Pop(weight).(int)
			heap.Push(weight, course[0])
		}
	}

	return taken
}
