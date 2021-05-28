package greedy

/*
Question 1695: https://leetcode.com/problems/maximum-erasure-value/

Idea:
- Sliding windows
- Use either a hashmap or a lastOccurence array to check if a num exist in current sum

Test case:
[4,2,4,5,6]
[5,2,2,2,2,3,1,2,5]
[558,508,782,32,187,103,370,607,619,267,984,10]
**/

func maximumUniqueSubarray(nums []int) int {
	var max, cur, last int
	pos := make(map[int]int)
	for i, val := range nums {
		if _, ok := pos[val]; ok {
			holder := pos[val]
			for j := last; j <= pos[val]; j++ {
				cur -= nums[j]
				delete(pos, nums[j])
			}
			last = holder + 1
		}
		cur += val
		pos[val] = i
		if cur > max {
			max = cur
		}
	}
	return max
}
