package constructive

/*
Question 118: https://leetcode.com/problems/pascals-triangle/

**/

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		cur := make([]int, i+1)
		for j := range cur {
			if j == 0 || j == len(cur)-1 {
				cur[j] = 1
			} else {
				cur[j] = res[i-1][j] + res[i-1][j-1]
			}
		}
		res[i] = cur
	}
	return res
}
