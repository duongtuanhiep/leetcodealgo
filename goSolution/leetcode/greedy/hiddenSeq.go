package greedy

/*

Question 2145: https://leetcode.com/problems/count-the-hidden-sequences/

Greedy: Keep track of both upper and lower bound for both hi,lo
Iterate through arrays and return diff.
Keys takeaway is that the differences must be correct so it only allow a range from lo-hi and not all possible combination.
**/

func numberOfArrays(differences []int, lower int, upper int) int {
	lo, hi := lower, upper
	for _, dif := range differences {
		lo += dif
		hi += dif

		if lo < lower {
			lo = lower
		}
		if hi > upper {
			hi = upper
		}
		if lo > hi {
			return 0
		}
	}

	return hi - lo + 1
}
