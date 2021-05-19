package constructive

/*
Question 1862: https://leetcode.com/problems/sum-of-floored-pairs/

Idea:
- Storing the frequency.

- Optimizing the loop.
**/

func sumOfFlooredPairs(nums []int) int {
	var res int

	count := make([]int, 200001)
	checked := make(map[int]bool)
	for _, val := range nums {
		count[val]++
	}

	for i := len(count) - 2; i >= 0; i-- {
		count[i] += count[i+1]
	}

	for _, num := range nums {
		if !checked[num] {
			for step := 1; (step+1)*num < len(count) && count[step*num] != 0; step++ {
				res += (count[num] - count[num+1]) * step * (count[step*num] - count[(step+1)*num])
			}
		}
		checked[num] = true
	}

	return res % (1000000000 + 7)
}
