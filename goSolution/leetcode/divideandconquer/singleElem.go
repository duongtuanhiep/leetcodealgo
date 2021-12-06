package divideandconquer

/*
Question 540: https://leetcode.com/problems/single-element-in-a-sorted-array/

Observation: since all elements appear twice and required elem appears one,
total number len(array) will always be odds.
Using binary search, we can figure out location of elems are to the left or right
of current search elem

Some test case
[1,1,2,3,3,4,4,8,8]
[3,3,7,7,10,11,11]
[1,1,3,3,4,4,5,5,7]
**/
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] != nums[mid+1] {
			high = mid
		} else {
			low = mid + 2
		}
	}
	return nums[low]
}
