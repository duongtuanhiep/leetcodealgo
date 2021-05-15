package greedy

/*
Question 1855: https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/

Greedy: Try increasing index in nums2 as much as possible, if not possible increase index
in nums1 so that it satify the initial conditions. Keep track of the distance.

**/
func maxDistance(nums1 []int, nums2 []int) int {
	var max, cur = 0, 0
	var i, j = 0, 0
	for ; i < len(nums1) && j < len(nums2); j++ {
		cur = j - i
		for i < len(nums1) && nums1[i] > nums2[j] {
			i++
			cur--
		}
		if cur > max {
			max = cur
		}
	}

	return max
}
