package greedy

import "sort"

/*
Question 1433: https://leetcode.com/problems/check-if-a-string-can-break-another-string/

Observation: A string can break if any int(charA) >= int(charB)
Suggesting that if there exists an order where any i-th largest of A should be larger
than that of B there will be a case where A would break B.

Solution: Sort both s1 and s2. Have "left" and "right" denotes the number of times
when string s1 and string s2 can "break" each other at a given moment when iterating.
Iterate through the len of one string, calculate the number of left and right and at the
end compares both. If either left or right == len(string) then we know that there will
be a case where s1 would break s2 on any position and vice-versa.
**/

func checkIfCanBreak(s1 string, s2 string) bool {
	var left, right int
	rune1 := []rune(s1)
	rune2 := []rune(s2)
	sort.SliceStable(rune1, func(i, j int) bool { return rune1[i] < rune1[j] })
	sort.SliceStable(rune2, func(i, j int) bool { return rune2[i] < rune2[j] })
	for i := 0; i < len(rune1); i++ {
		if rune1[i] >= rune2[i] {
			left++
		}
		if rune1[i] <= rune2[i] {
			right++
		}
	}
	if left == len(rune1) || right == len(rune1) {
		return true
	}
	return false
}
