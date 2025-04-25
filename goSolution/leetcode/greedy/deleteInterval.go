package greedy

import (
	"sort"
)

/*

Question 435: https://leetcode.com/problems/non-overlapping-intervals

Greedy:
- Sort by end points
- If the previous end is touch, we delete the current one.
- This way we can keep the end as "early" as possible.

**/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res, prevEnd := 0, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if prevEnd > intervals[i][0] {
			res++
		} else {
			prevEnd = intervals[i][1]
		}
	}

	return res
}
