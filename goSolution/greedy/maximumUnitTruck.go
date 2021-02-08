package greedy

import "sort"

/*
Question 1710: https://leetcode.com/problems/maximum-units-on-a-truck/

Idea: Put all the box of most weight type available until full

**/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.SliceStable(boxTypes, func(a, b int) bool { return boxTypes[a][1] < boxTypes[b][1] })
	var res int
	for truckSize > 0 {
		if boxTypes[0][0] == 0 {
			boxTypes = boxTypes[1:]
		}
		res += boxTypes[0][1]
		boxTypes[0][0]--
		truckSize--
	}
	return res
}
