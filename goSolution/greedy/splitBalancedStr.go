package greedy

// Question 1221
// https://leetcode.com/problems/split-a-string-in-balanced-strings/

/*
- The initial strings is a balanced string of "L" and "R".
- Can iterate through the inital strings, denoting as L as -1 and R as +1,
any balanced "&L&R&" strings will satify len(str) > 0 and sum as 0
- Iterate through the beginning strings, increase count whenever see the
curSum == 0
**/

func balancedStringSplit(s string) int {
	var totalCount int
	var curSum int
	for _, cur := range s {
		if cur == rune('L') {
			curSum--
		} else {
			curSum++
		}
		if curSum == 0 {
			totalCount++
		}
	}
	return totalCount
}
