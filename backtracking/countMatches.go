package backtracking

/*
Question 1688: https://leetcode.com/problems/count-of-matches-in-tournament/


**/

func numberOfMatches(n int) int {
	if n < 2 {
		return 0
	}
	if n%2 == 0 {
		return n/2 + numberOfMatches(n/2)
	}
	return (n-1)/2 + numberOfMatches((n-1)/2+1)
}
