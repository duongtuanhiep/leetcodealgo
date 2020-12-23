package sorting

import "sort"

// Solution: Brute-force
// double checks if each of any abitrary interval is overlap. merge those into one, continue with the newly merged array
func merge(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	var res [][]int
	res = append(res, intervals[0])
	for _, valPair := range intervals {
		for _, valRes := range res {
			if valRes[0] <= valPair[0] && valPair[0] <= valRes[1] {
				if valPair[1] > valRes[1] {
					valRes[1] = valPair[1]
				}
			} else if valRes[0] <= valPair[1] && valPair[1] <= valRes[1] {
				if valPair[0] < valRes[0] {
					valRes[0] = valPair[0]
				}
			} else if valPair[0] <= valRes[0] && valPair[1] >= valRes[0] {
				valRes[0] = valPair[0]
				valRes[1] = valPair[0]
			} else if valPair[0] > valRes[1] || valPair[1] < valRes[0] {
				res = append(res, valPair)
			}
			break
		}
	}
	return res
}
