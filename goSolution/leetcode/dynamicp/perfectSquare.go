package dynamicp

import "math"

/*
Question 279: https://leetcode.com/problems/perfect-squares/

1
2
3
4
5
6
7
8
9
10
**/

func numSquares(n int) int {
	arr := make([]int, n+1)
	for i := 1; i < len(arr); i++ {
		if isSqrt(i) {
			arr[i] = 1
		} else {
			minNum := arr[1] + arr[i-1]
			for j := i - 1; j > 0; j-- {
				minNum = min(minNum, arr[i-j]+arr[j])
			}
			arr[i] = minNum
		}
	}

	return arr[n]
}

func isSqrt(n int) bool {
	return math.Sqrt(float64(n)) == float64(int(math.Sqrt(float64(n))))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
