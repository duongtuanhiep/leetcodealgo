package dfs

import (
	"math"
	"strconv"
)

/*

Question 1415: https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/

Obascervation:
- Instead of dealing with string, we can try dealing with number since a,b,c is practically 1,2,3 in lexicographical terms
- Generate all number of length n via DFS.
- Since it DFS, we can already return k as it is the index of kth elem from arr.
**/

func getHappyString(n int, k int) string {
	available := []int{1, 2, 3}
	allCombo := generateHappy(n, -1, available)
	if k > len(allCombo) {
		return ""
	}

	var res string
	for _, char := range strconv.Itoa(allCombo[k-1]) {
		switch char {
		case '1':
			res += "a"
		case '2':
			res += "b"
		case '3':
			res += "c"
		}
	}

	return res
}

func generateHappy(length, last int, available []int) []int {
	if length == 0 {
		return nil
	}
	var res []int

	currentNumPrefix := int(math.Pow(10.0, float64(length-1)))
	for _, val := range available {
		if val == last {
			continue
		}

		combinations := generateHappy(length-1, val, available)
		// case when this is not the last number
		for _, c := range combinations {
			combined := val*currentNumPrefix + c
			res = append(res, combined)
		}

		// case when this is already the last number
		if len(combinations) == 0 {
			res = append(res, val)
		}
	}

	return res
}
