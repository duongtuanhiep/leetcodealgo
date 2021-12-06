package wc242

import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	if len(dist) < 2 {
		if float64(dist[0]) < hour {
			return dist[0]
		}
		return -1
	}

	sumDist := 0
	for _, val := range dist {
		sumDist += val
	}

	var res int = 1
	var totalTime float64
	for _, val := range dist {
		var potionTime float64 = (hour - totalTime) / float64(sumDist) * float64(val)

		minCurSpeed := roundUp(float64(val) / potionTime)
		totalTime += float64(val) / roundUp(minCurSpeed)

		if int(minCurSpeed) > res {
			res = int(minCurSpeed)
		}
		if totalTime > hour {
			return -1
		}

		sumDist -= val
		totalTime = roundUp(totalTime)
	}

	return res
}

func roundUp(num float64) float64 {
	if math.Round(num) < num {
		return (math.Round(num) + 1)
	}
	return (math.Round(num))
}
