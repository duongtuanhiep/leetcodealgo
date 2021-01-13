package greedy

import "sort"

/*
Question 881: https://leetcode.com/problems/boats-to-save-people/

Idea:
- Creates map of k to keep track of num

**/

// Brute Force: O(N^2)

func numRescueBoats(people []int, limit int) int {

	var result int
	sort.Ints(people)
	var i int
	for j := len(people) - 1; j >= 0; j-- {
		if people[j]+people[i] <= limit {
			i++
			result++
		} else {
			result++
		}

		if i > j {
			break
		}
	}
	return result
}

/*
Intuition:
- Pairing the lightest with the heaviest possible.
- Sort the array, creates 2 pointers.
**/
func numRescueBoats(people []int, limit int) int {
	var result int
	sort.Ints(people)
	var i int
	for j := len(people) - 1; j >= i; j-- {
		if people[j]+people[i] <= limit {
			i++
		}
		result++
	}
	return result
}

// Improvement: Bucket sort since individual can always be carried by a boat
// https://leetcode.com/problems/boats-to-save-people/discuss/197063/easy-thought-process-to-improve-from-O(nlogn)-to-O(n)
