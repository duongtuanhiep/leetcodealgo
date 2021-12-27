package bitmanipulation

import "math/bits"

/*
Question 476: https://leetcode.com/problems/number-complement/


right shift and AND 1, if == 1 means last num is 1
**/

func findComplement(num int) int {
	maxInt := 9223372036854775807
	shift := bits.LeadingZeros64(uint64(num))
	return (maxInt >> shift) &^ num
}
