package array

/*
Question 80: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

Approach: 2 pointers
We try to keep track of the current and the current good portion of the array. We
swap on each valid number (when its different from both last 2 good) on each iteration.

By default, the first 2 is always "good"

TEST:
[1,1,1,2,2,3]
[0,0,0,1,1,1,1,2,2,2,3,3,3,4]
[1,1]
[1]
**/

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var pos int
	for i := range nums {
		if i < 2 {
			pos++
			continue
		}

		if nums[i] != nums[pos-2] || nums[i] != nums[pos-1] {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}

	return pos
}
