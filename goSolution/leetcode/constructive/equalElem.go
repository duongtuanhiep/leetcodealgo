package constructive

import "math/rand"

/*
Question 462: https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/

*/

func minMoves2(nums []int) int {
	var median int
	if len(nums)%2 == 1 {
		median = findKthLargest(nums, len(nums)/2+1)
	} else {
		median = (findKthLargest(nums, len(nums)/2+1) + findKthLargest(nums, len(nums)/2+1)) / 2
	}

	var res int
	for _, val := range nums {
		res += abs(val, median)
	}

	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func findKthLargest(nums []int, k int) int {
	return findKthElem(nums, len(nums)-1, 0, len(nums)-k)
}

// Observation on boundary
// - kth largest element, or len() - k smallest elem, meaning the index len() - k  in sorted array
// Selection algo, on memory, k = index of kth largest indexes (len-k) , lo = first index, hi = last index of processing arr
func findKthElem(nums []int, hi, lo, k int) int {

	// Select Pivot
	pivot := lo + rand.Intn(hi-lo+1)

	// Swap to front of array
	nums[pivot], nums[lo] = nums[lo], nums[pivot]

	// Perform partrition
	i := lo + 1
	for j := lo + 1; j <= hi; j++ {
		if nums[j] < nums[lo] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	// Swap back to place, check location relatively to k
	nums[lo], nums[i-1] = nums[i-1], nums[lo]

	// Compare pivot, if lucky happy
	if i-1 == k {
		return nums[i-1]
	}

	// Recurse on either left or right side with respective boundary
	// case when k > i-1
	if k > i-1 {
		return findKthElem(nums, hi, i, k)
	}

	// case when k < i-1
	return findKthElem(nums, i-2, lo, k)
}
