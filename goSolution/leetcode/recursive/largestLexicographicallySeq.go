package recursivealgo

import "fmt"

/*

Question 1718: https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/

Some observation:
- Len is always 2n - 1

Brute-force:
- Just try to fill in number as large as possible first.

Optimization:
- Instead of filling 1 number at a time we can try to fill in a whole pair since the latter number's index is always bounded

**/

func constructDistancedSequence(n int) []int {

	// filling remaining number
	remaining := make([]int, n)
	for i := range remaining {
		remaining[i] = 1
	}

	current := make([]int, n*2-1)
	legit := fill(current, remaining)
	fmt.Println(legit)
	return current
}

func fill(current, remaining []int) bool {
	nextIndex := 0

	for nextIndex < len(current) && current[nextIndex] != 0 {
		nextIndex++
	}

	if nextIndex >= len(current) {
		return true
	}

	// try to fill the largest number first
	for i := len(remaining) - 1; i >= 0; i-- {
		number := i + 1

		if remaining[i] > 0 {
			// if number != 1, put the remaining of the pair into a valid position ahead in the arr
			if number != 1 {
				// cant put it anywhere else anymore
				if nextIndex+number >= len(current) {
					return false
				}

				// check if possible to put number here
				if current[nextIndex+number] != 0 {
					continue
				}
				current[nextIndex+number] = number
			}

			// fill current pos and mark as used
			current[nextIndex] = number
			remaining[i] = 0

			possible := fill(current, remaining)
			// revert the marking in case it's invalid later on
			if !possible {
				if number > 1 {
					current[nextIndex+number] = 0
				}
				current[nextIndex] = 0
				remaining[i] = 1 // put back to remaining
			} else if possible {
				return true
			}
		}
	}

	return false
}
