package greedy

// https://leetcode.com/problems/last-stone-weight/

/* Naive implementation:
- Run through the array , get 2 largest elem,
- Swap those to the end of array and trim off those 2 element
- Added in result of max - secondMax if > 0
- continue until array has len == 1

*/
func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		var first, second, firstPos, secondPos int
		for i := range stones {
			if stones[i] >= first {
				first, second = stones[i], first
				firstPos, secondPos = i, firstPos
			} else if stones[i] >= second {
				second = stones[i]
				secondPos = i
			}
		}

		// Swapping to back of array, trim
		if firstPos == len(stones)-1 {
			stones[secondPos], stones[len(stones)-2] = stones[len(stones)-2], stones[secondPos]
			stones = stones[:len(stones)-2]
		} else if secondPos == len(stones)-1 {
			stones[firstPos], stones[len(stones)-2] = stones[len(stones)-2], stones[firstPos]
			stones = stones[:len(stones)-2]
		} else {
			stones[firstPos], stones[len(stones)-1] = stones[len(stones)-1], stones[firstPos]
			stones[secondPos], stones[len(stones)-2] = stones[len(stones)-2], stones[secondPos]
			stones = stones[:len(stones)-2]
		}

		// trim arr

		if first-second > 0 {
			stones = append(stones, first-second)
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
