package array

/*

Question 873: https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/

**/

func lenLongestFibSubseq(arr []int) int {
	exist := make(map[int]struct{})

	for _, val := range arr {
		exist[val] = struct{}{}
	}

	var res int
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			// At least found a fib
			if _, found := exist[arr[i]+arr[j]]; found {
				first, second, cur := arr[j], arr[i]+arr[j], 3
				for _, found := exist[first+second]; found; {
					cur++
					first, second = second, first+second
					_, found = exist[first+second]
				}
				if cur > res {
					res = cur
				}
			}
		}
	}
	return res
}
