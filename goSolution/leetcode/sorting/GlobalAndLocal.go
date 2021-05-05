package sorting

/*
Question 775: https://leetcode.com/problems/global-and-local-inversions/

Idea:
- We can scan through the array for local inversion since it is strickly A[i] > A[i+1] with 0 < i < N

**/

func isIdealPermutation(A []int) bool {
	var local int
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			local++
		}
	}

	var global int
	for i := 0; i < len(A); i++ {
		if i != A[i] {
			counter := 0
			for j := i + 1; j < len(A); j++ {
				if A[i] > A[j] {
					counter++
				}
			}
			global += counter
		}
	}

	return global == local
}
