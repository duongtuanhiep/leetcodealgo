package array

// Dumb O(n^2)
func maxScoreSightseeingPair(A []int) int {
	var cur, max int
	front, back := A, A
	for pos := range front {
		front[pos] += pos
	}
	for pos := range back {
		back[pos] -= pos
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			cur = front[i] + back[j]
			if cur > max {
				max = cur
			}
		}
	}
	return max
}
