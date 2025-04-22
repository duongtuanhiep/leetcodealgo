package divideandconquer

import "slices"

/*

Question 2563:https://leetcode.com/problems/count-the-number-of-fair-pairs

**/

func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	var res int64
	for i := range nums {
		hiTarget, loTarget := upper-nums[i], lower-nums[i]
		hi, lo := bsRightMostEqualOrSmaller(nums, i+1, hiTarget), bsLeftMostEqualOrGreater(nums, i+1, loTarget)
		if lo <= hi && lo != -1 && hi != -1 {
			res += int64(hi-lo) + 1
		}
	}

	return res
}

func bsLeftMostEqualOrGreater(nums []int, i, target int) int {
	hi, lo := len(nums)-1, i
	targetIndex := -1

	for lo <= hi {
		mid := (hi + lo) / 2
		if target <= nums[mid] {
			targetIndex = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return targetIndex
}

func bsRightMostEqualOrSmaller(nums []int, i, target int) int {
	hi, lo := len(nums)-1, i
	targetIndex := -1

	for lo <= hi {
		mid := (hi + lo) / 2
		if target >= nums[mid] {
			targetIndex = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return targetIndex
}
