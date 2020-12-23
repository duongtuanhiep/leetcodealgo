package divideandconquer

import (
	"sort"
)

// Sorting solution. Sort-based. Runtime : O(nLogn)
func findKthLargest(nums []int, k int) int {

	sort.Ints(nums)
	var i int = 1 // current largest elem, representing at pos
	for pos := len(nums) - 1; pos > 0; pos-- {
		if nums[pos] > nums[pos-1] {
			i++
			if i == k {
				return nums[pos-1]
			}
		}
	}
	return 0
}
