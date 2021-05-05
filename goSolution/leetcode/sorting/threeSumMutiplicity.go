package sorting

/*
Question 923: https://leetcode.com/problems/3sum-with-multiplicity/

Idea:
- Do a bucket sort since the constraint allow you to do so efficiently.
- We then try to optimize inner step to cut down the constant.

**/

func threeSumMulti(arr []int, target int) int {
	var res int
	bucket := make([]int, 101)

	for _, val := range arr {
		bucket[val]++
	}

	j := 100
	for i := 0; i < len(bucket); i++ {
		visited := make(map[int]bool)
		if i > j {
			break
		}
		if target-i-j < 0 {
			j--
		}
		if bucket[i] > 0 && bucket[j] > 0 && bucket[target-i-j] > 0 {

		}
	}

	return res
}
