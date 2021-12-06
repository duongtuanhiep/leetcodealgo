package constructive

/*
Question 556: https://leetcode.com/problems/reshape-the-matrix/


**/

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) == r*c {
		newMat := make([][]int, r)
		holder := []int{}
		for i := range mat {
			for _, val := range mat[i] {
				holder = append(holder, val)
			}
		}

		counter := 0
		for i := range newMat {
			newMat[i] = holder[counter : counter+c]
			counter += c
		}

		return newMat
	}
	return mat
}
