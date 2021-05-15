package array

/*
Question 204: https://leetcode.com/problems/count-primes/solution/

Sieve of Eratosthenes: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
**/

func countPrimes(n int) int {
	nonPrime := make([]bool, n+1)
	var res int
	for i := 2; i < len(nonPrime)-1; i++ {
		if !nonPrime[i] {
			for j := 2; j*i < len(nonPrime)-1; j++ {
				nonPrime[j*i] = true
			}
			res++
		}
	}
	return res
}
