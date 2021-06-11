package backtracking

import "fmt"

/*
Question 1447: https://leetcode.com/problems/simplified-fractions/

*/

func simplifiedFractions(n int) []string {
	exist := make(map[float64]bool)
	var res []string
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			if !exist[float64(j)/float64(i)] {
				res = append(res, fmt.Sprintf("%d/%d", j, i))
				exist[float64(j)/float64(i)] = true
			}
		}
	}
	return res
}
