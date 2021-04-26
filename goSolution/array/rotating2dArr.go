package array

/*
Question 48: https://leetcode.com/problems/rotate-image/

Idea:
Switching 1 element results in swapping values at 4 element in 4 "edge" of the matrix.
We can derive an algorithm from a few plotted run

Imagine a n = 4 len matrices. Each iteration it will swap 4 elements.
	- First iteration
		[X - i - i - X]
		[i - i - i - i]
		[i - i - i - i]
		[X - i - i - X]
	- Second iteration
		[i - X - i - i]
		[i - i - i - X]
		[X - i - i - i]
		[i - i - X - i]
	...


Example test case:
[[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]
[[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]
**/
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < n-i; j++ {
			matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i], matrix[i][j] = matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i]
		}
	}
}
