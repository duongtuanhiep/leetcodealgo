package bitmanipulation

import "math"

/*

Question 3513: https://leetcode.com/problems/number-of-unique-xor-triplets-i/

Fucking magician math bullshit and not a fucking "algorithm" questions:
- The whole algorithm is devised from proving that from an arrays of [1,n], you can XOR them and it is bounded by
the maximum number of bits represented (max bitslength)
- The only ways you'd devise this during interview settings is to dry a few cases and try to connect the dot. Can you prove correctness ? probably not.
**/

func uniqueXorTriplets(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var bitCount int
	for i := len(nums); i > 0; {
		bitCount++
		i >>= 1
	}

	return (int(math.Pow(2.0, float64(bitCount))))
}
