package dynamicp

/*
Question 70: https://leetcode.com/problems/climbing-stairs/

Idea:
- Recursion relation: Every step i can be climb to from step i-1 and i-2 ->
total possible way to step to is the sum of step at i-1 + i -2
- We then can store it to the array so that we dont have to recalculate each
time over and over again.
- Base case: step = 0 : 1 way, step = 1: 1 way
**/

func climbStairs(n int) int {
	steps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i < 2 {
			steps[i] = 1
		} else {
			steps[i] = steps[i-1] + steps[i-2]
		}
	}

	return steps[n]
}
