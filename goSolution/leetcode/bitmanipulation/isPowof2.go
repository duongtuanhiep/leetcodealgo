package bitmanipulation

/*
Question 231: https://leetcode.com/problems/power-of-two/

return (n-1 AND n) == 0
**/

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n>>1*2 == n {
		return isPowerOfTwo(n >> 1)
	}
	return false
}
