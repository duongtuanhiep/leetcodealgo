package dynamicp

/*
Question 304: https://leetcode.com/problems/range-sum-query-2d-immutable/

**/

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

type NumMatrix struct {
	Arr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var m [][]int
	for i := 0; i < len(matrix); i++ {
		m = append(m, make([]int, len(matrix[0])))
	}

	for i := 0; i < len(m); i++ {
		if i == 0 {
			m[i][0] = matrix[i][0]
		} else {
			m[i][0] = m[i-1][0] + matrix[i][0]
		}
	}

	for i := 0; i < len(m[0]); i++ {
		if i == 0 {
			m[0][i] = matrix[0][i]
		} else {
			m[0][i] = m[0][i-1] + matrix[0][i]
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			m[i][j] = m[i-1][j] + m[i][j-1] - m[i-1][j-1] + matrix[i][j]
		}
	}

	// fmt.Println(m)
	return NumMatrix{m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var res = this.Arr[row2][col2]
	if col1 > 0 {
		res -= this.Arr[row2][col1-1]
	}
	if row1 > 0 {
		res -= this.Arr[row1-1][col2]
	}
	if row1 > 0 && col1 > 0 {
		res += this.Arr[row1-1][col1-1]
	}

	return res
}
