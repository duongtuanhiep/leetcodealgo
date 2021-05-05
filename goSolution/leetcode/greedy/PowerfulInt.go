package greedy

/*
Question 970: https://leetcode.com/problems/powerful-integers/

**/
func powerfulIntegers(x int, y int, bound int) []int {

	seen := make(map[int]bool)
	res := []int{}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if pow(x, i)+pow(y, j) <= bound {
				seen[pow(x, i)+pow(y, j)] = true
			} else {
				break
			}
		}
		if pow(x, i) > bound {
			break
		}
	}
	for k := range seen {
		res = append(res, k)
	}
	return res
}

func pow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}
