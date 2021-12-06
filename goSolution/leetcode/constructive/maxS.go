package constructive

import "sort"

/*
Question 1465: https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/

Idea:
- calculate each square.
- or maybe just calculate the largest one.
- or bucket sort.
- or init 2 heaps

5
4
[1,2,4]
[1,3]
5
4
[3,1]
[1]
6
2
[5]
[1]
**/

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	var height, width int
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	for i := range horizontalCuts {
		if i == 0 || i == len(horizontalCuts)-1 {
			height = max(height, max(horizontalCuts[0], h-horizontalCuts[len(horizontalCuts)-1]))
		}
		if i > 0 && height < horizontalCuts[i]-horizontalCuts[i-1] {
			height = horizontalCuts[i] - horizontalCuts[i-1]
		}
	}
	for i := range verticalCuts {
		if i == 0 || i == len(verticalCuts)-1 {
			width = max(width, max(verticalCuts[0], w-verticalCuts[len(verticalCuts)-1]))
		}
		if i > 0 && width < verticalCuts[i]-verticalCuts[i-1] {
			width = verticalCuts[i] - verticalCuts[i-1]
		}
	}

	return (height * width) % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
