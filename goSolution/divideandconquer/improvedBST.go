package divideandconquer

/*
Question 34: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

Solution: https://en.wikipedia.org/wiki/Binary_search_algorithm
**/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	a, b := bsLeftmost(nums, target), bsRightmost(nums, target)
	if a >= len(nums) || b < 0 || nums[a] != nums[b] || nums[a] != target {
		return []int{-1, -1}
	}
	return []int{a, b}
}

func bsLeftmost(nums []int, target int) int {
	hi, lo := len(nums), 0
	for lo < hi {
		m := (hi + lo) / 2
		if nums[m] < target {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return lo
}

func bsRightmost(nums []int, target int) int {
	hi, lo := len(nums), 0
	for lo < hi {
		m := (hi + lo) / 2
		if nums[m] > target {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return hi - 1
}
