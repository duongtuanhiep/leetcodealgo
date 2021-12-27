package greedy

import "sort"

/*
Question 56: https://leetcode.com/problems/merge-intervals/
**/

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(a, b int) bool { return intervals[a][0] < intervals[b][0] })
	cur := intervals[0]
	for _, pair := range intervals {
		if cur[1] < pair[0] {
			res = append(res, cur)
			cur = pair
		} else if cur[1] >= pair[0] && cur[1] < pair[1] {
			cur[1] = pair[1]
		}
	}
	return append(res, cur)
}
