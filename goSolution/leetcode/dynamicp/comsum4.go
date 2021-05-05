package dynamicp

/*
Question 377: https://leetcode.com/problems/combination-sum-iv/


Idea: dynamic programming with memoization
- for any number i and current number j, we can "reach" i from every pos[i-j] with i-j > 0
- Recursion : S[i] = SUM(S[i-j])
- Base case: Every number i = j can be reach from -> S[i-j] = S[0] (which is 1)
- We create an array to store all the middle steps so that we dont have to do that again
*/

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
