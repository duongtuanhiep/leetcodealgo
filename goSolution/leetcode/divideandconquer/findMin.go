package divideandconquer

import "fmt"

/*

Question 153: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array
Question 33: https://leetcode.com/problems/search-in-rotated-sorted-array

Binary search.
**/

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	res := -5001
	for lo <= hi {
		mid := (hi + lo) / 2
		if nums[mid] >= nums[hi] && nums[mid] >= nums[lo] { // Arrays been rotated past mid point
			lo = mid + 1
		} else {
			hi = mid
		}

		if res > nums[mid] || res == -5001 {
			res = nums[mid]
		}
	}

	return res
}

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (hi + lo) / 2
		if nums[mid] == target {
			return mid
		}

		// Determine if target is in sorted or un-sorted
		fmt.Println(nums[lo], nums[mid], nums[hi])
		if nums[lo] <= nums[mid] { // When the section is sorted
			if target < nums[mid] && target >= nums[lo] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // When the section is un-sorted
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
