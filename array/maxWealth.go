package array

func maximumWealth(accounts [][]int) int {
	var max int
	for i := range accounts {
		var cur int
		for j := range accounts[i] {
			cur += accounts[i][j]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
