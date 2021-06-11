package bc53

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	var res, cur int
	for i := 0; i < len(nums)/2; i++ {
		cur = nums[i] + nums[len(nums)-1-i]
		if cur > res {
			res = cur
		}
	}

	return res
}
