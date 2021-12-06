package bc53

import "sort"

func getBiggestThree(grid [][]int) []int {
	var resArr []int
	for i := range grid {
		for j := range grid[i] {
			edge := checkedge(grid, i, j)
			if edge {
				resArr = append(resArr, grid[i][j])
			} else {
				countStep := getMaxDimension(grid, i, j)
				for count := 0; count <= countStep; count++ {
					var cur int
					for k := 0; k < count; k++ {
						cur += grid[i-count+k][j+k] + grid[i+k][j+count-k] + grid[i+count-k][j-k] + grid[i-k][j-count+k]
					}
					resArr = append(resArr, cur)
				}
			}
		}
	}

	var res []int
	sort.Ints(resArr)
	last := -1
	for i := len(resArr) - 1; i >= 0; i-- {
		if resArr[i] != last {
			res = append(res, resArr[i])
			last = resArr[i]
		}
		if len(res) == 3 {
			break
		}
	}

	return res
}

func checkedge(grid [][]int, i, j int) bool {
	if i == 0 || j == 0 || len(grid)-1 == i || len(grid[0])-1 == j {
		return true
	}
	return false
}

func getMaxDimension(grid [][]int, i, j int) int {
	a := min(len(grid)-1-i, i)
	b := min(len(grid[0])-1-j, j)
	if a < 0 {
		a = 0
	}
	if b < 0 {
		b = 0
	}
	return min(a, b)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
