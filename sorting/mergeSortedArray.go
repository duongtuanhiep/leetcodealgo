package sorting

/*
Question 88: https://leetcode.com/problems/merge-sorted-array/

Idea:
	- Retain array nums1, make a copy called holder and use holder
	instead of nums1. Apply normal merging technique.
**/

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, len(nums1))
	copy(res, nums1)
	var i, j = 0, 0
	for i < m || j < n {
		if i < m && j < n {
			if res[i] < nums2[j] {
				nums1[i+j] = res[i]
				i++
			} else {
				nums1[i+j] = nums2[j]
				j++
			}
		} else if i == m {
			nums1[i+j] = nums2[j]
			j++
		} else if j == n {
			nums1[i+j] = res[i]
			i++
		}
	}
}

// Make a queue to handle merge
func merge(nums1 []int, m int, nums2 []int, n int) {
	// var j int = 0
	var queue []int
	var j int = 0
	if len(nums2) == 0 {
		return
	}

	for i := range nums1 {
		if len(queue) == 0 {
			if nums1[i] > nums2[j] {
				queue = append(queue, nums1[i])
				nums1[i] = nums2[j]
				j++
			} else if i >= m {
				nums1[i] = nums2[j]
				j++
			}
		} else if j < n {
			if queue[0] < nums2[j] && i < m+j {
				queue = append(queue, nums1[i])
				nums1[i] = queue[0]
				queue = queue[1:]
			} else {
				queue = append(queue, nums1[i])
				nums1[i] = nums2[j]
				j++
			}
		} else {
			queue = append(queue, nums1[i])
			nums1[i] = queue[0]
			queue = queue[1:]
		}
	}
}
