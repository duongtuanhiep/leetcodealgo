package dynamicp

/*
Question 120: https://leetcode.com/problems/triangle/

Idea:
- We can try to calculate what is the min value cost at position i from position i+1 in array triangle.


*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 1 {
		return triangle[0][0]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j+1]
			} else {
				triangle[i][j] += triangle[i+1][j]
			}
		}
	}

	return triangle[0][0]
}
