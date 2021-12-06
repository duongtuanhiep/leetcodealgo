package constructive

/*
Question 54: https://leetcode.com/problems/spiral-matrix/

[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
[[1,2,3,4],[5,6,7,8]]
[[1,2,3,4],[5,6,7,8],[9,10,11,12],[13,14,15,16]]

**/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 1 {
		return matrix[0]
	}
	var res []int
	var k, lim int
	i, j := len(matrix)-1, len(matrix[0])-1
	if len(matrix) > len(matrix[0]) {
		lim = len(matrix[0]) / 2
	} else {
		lim = len(matrix) / 2
	}
	for k <= lim {
		for s := k; s < j; s++ {
			res = append(res, matrix[k][s])
		}
		for s := k; s < i; s++ {
			res = append(res, matrix[s][j-k])
		}
		for s := j - k; s > 0; s-- {
			res = append(res, matrix[i-k][s])
		}
		for s := i - k; s > 0; s-- {
			res = append(res, matrix[s][k])
		}
		k++
	}
	return res[:len(matrix)*len(matrix[0])]
}
