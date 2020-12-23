package dynamicp

import "math"

func maxSubArray(nums []int) int {

	var maxGlobal, maxLocal int

	//base case
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxGlobal = nums[0]
	maxLocal = nums[0]
	for i := 1; i < len(nums); i++ {
		maxLocal = int(math.Max(float64(maxLocal+nums[i]), float64(nums[i])))
		if maxLocal > maxGlobal {
			maxGlobal = maxLocal
		}
	}

	return maxGlobal
}
