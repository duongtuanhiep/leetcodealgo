package selection

/*
Question 1870: https://leetcode.com/problems/minimum-speed-to-arrive-on-time/submissions/


Idea: Binary search for the desired integer speed since it is within the range [0,1e7]
**/

import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	var lo, hi int = 0, 1e7 + 1
	for lo < hi {
		var timeNeed float64
		pivot := (hi + lo) / 2
		for _, val := range dist[:len(dist)-1] {
			timeNeed += roundUp(float64(val) / float64(pivot))
		}
		timeNeed += float64(dist[len(dist)-1]) / float64(pivot)
		if timeNeed > hour {
			lo = pivot + 1
		} else {
			hi = pivot
		}
	}
	if lo > 1e7 {
		return -1
	}
	return lo
}

func roundUp(num float64) float64 {
	if math.Round(num) < num {
		return math.Round(num) + 1
	}
	return math.Round(num)
}
