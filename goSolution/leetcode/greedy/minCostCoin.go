package greedy

/*
Question 1217: https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
**/

func minCostToMoveChips(position []int) int {
	var even, odd int
	for _, pos := range position {
		if pos%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		return odd
	}
	return even
}
