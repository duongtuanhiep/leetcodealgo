package greedy

import "sort"

/*
Question 452: https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/


Idea: Sort the balloons by starting points. On each iteration, if you have to place
"arrow" try to place it as further back as possible
**/

func findMinArrowShots(points [][]int) int {
	var res, front, back int
	sort.Slice(points, func(a, b int) bool { return points[a][0] < points[b][0] })
	front, back = -9223372036854775808, -9223372036854775808
	for _, pair := range points {
		if back < pair[0] {
			res++
			front, back = pair[0], pair[1]
		} else {
			if back >= pair[1] {
				back = pair[1]
			}
			if front <= pair[0] {
				front = pair[0]
			}
		}
	}
	return res
}
