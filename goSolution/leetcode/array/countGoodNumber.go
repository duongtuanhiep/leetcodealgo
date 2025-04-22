package array

/*
Question 1922: https://leetcode.com/problems/count-good-numbers/?envType=daily-question&envId=2025-04-13

// Have even/odd number with possible leading zeroes.
// We calculate possible combination can be made out of with even/odd

Need to use binary exponent: https://cp-algorithms.com/algebra/binary-exp.html
*
*/

func countGoodNumbers(n int64) int {
	mod := int64(1000000007)

	// use fast exponentiation to calculate x^y % mod
	quickmul := func(x, y int64) int64 {
		ret := int64(1)
		mul := x
		for y > 0 {

			// If bit is set in exponential, multiply it with res
			if y%2 == 1 {
				ret = ret * mul % mod
			}

			// Multiply exponent base
			mul = mul * mul % mod

			// Shift the bits to the right
			y = y >> 1
		}

		return ret
	}

	return int(quickmul(5, (n+1)/2) * quickmul(4, n/2) % mod)
}
