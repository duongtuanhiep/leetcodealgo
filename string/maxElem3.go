package string

import (
	"sort"
	"strconv"
)

/*
Question link : https://leetcode.com/problems/next-greater-element-iii/

Idea: Swapping any element x with y such that x > y

Solution O(N^2): Compare from right to left, first elem that is smaller, swap
	and then sort all the elements to the right of pivot
*/

// max 2147483647 int32
func nextGreaterElement(n int) int {
	numC := []rune(strconv.Itoa(n))
	var pivot int
	for i := len(numC) - 1; i >= 0; i-- {
		for j := len(numC) - 1; j > i; j-- {
			if numC[j] > numC[i] {
				pivot = i
				numC[i], numC[j] = numC[j], numC[i]
				goto here
			}
		}
	}

here:
	sortPls := numC[pivot+1 : len(numC)]
	sort.SliceStable(sortPls, func(i, j int) bool {
		return sortPls[i] < sortPls[j]
	})

	// Second Loop
	// Recreate from rune slice to int
	newResInt, _ := strconv.Atoi(string(append(numC[:pivot+1], sortPls...)))
	if newResInt > n && newResInt < 2147483647 {
		return newResInt
	}
	return -1
}
