package greedy

/*
Question 1413: https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/

Idea:
- Keep track of min val - cur sum
**/

func minStartValue(nums []int) int {
	var curSum, res int
	for _, val := range nums {
		if curSum+val < 0 {
			res -= (curSum + val)
			curSum = 0
		} else {
			curSum += val
		}
	}
	return res + 1
}
