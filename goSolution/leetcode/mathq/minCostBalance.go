package mathq

/*
Question 1551: https://leetcode.com/problems/minimum-operations-to-make-array-equal/submissions/

Self explanatory.
**/

func minOperations(n int) int {
	if n%2 == 0 {
		return (n / 2) * (n / 2)
	}
	return (n / 2) * (n/2 + 1)
}
