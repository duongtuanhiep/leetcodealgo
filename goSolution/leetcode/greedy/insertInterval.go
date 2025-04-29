package greedy

/*

Question 57: https://leetcode.com/problems/insert-interval

**/

func insert(intervals [][]int, newInterval []int) [][]int {
	finalInterval := make([][]int, 0, len(intervals)+1)

	// Binary search to slot newInterval into intervals
	inputIndex := bsInterval(intervals, newInterval[0])

	if inputIndex == -1 {
		intervals = append(intervals, newInterval)
	} else {
		intervals = append(intervals, nil) // make 1 more space
		copy(intervals[inputIndex+1:], intervals[inputIndex:])
		intervals[inputIndex] = newInterval
	}

	i := 0
	for i < len(intervals) {
		cur := intervals[i]
		j := i + 1
		for j < len(intervals) && cur[0] <= intervals[j][0] && cur[1] >= intervals[j][0] { // try extend forward
			cur[1] = max(cur[1], intervals[j][1])
			j++
		}
		finalInterval = append(finalInterval, cur)
		i = j
	}

	return finalInterval
}

func bsInterval(intervals [][]int, num int) int {
	lo, hi := 0, len(intervals)-1
	dest := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if intervals[mid][0] > num {
			dest = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return dest
}
