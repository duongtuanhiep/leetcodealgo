package selection

/*
Question 540: https://leetcode.com/problems/single-element-in-a-sorted-array/


**/

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var res = int
	lo, hi := 0, len(nums)-1
	for lo < hi {
		avg := (hi + lo) / 2
		adj := -1
		if avg > 0 && nums[avg-1] == nums[avg] {
			adj = avg - 1
		} else if avg < len(nums)-1 && nums[avg+1] == nums[avg] {
			adj = avg + 1
		} else {
			res = nums[avg]
			break
		}

		// recurse left 
		if adj < avg && (hi-avg)%2 != 0 || adj > avg
		
		// recurse right
	}

	return res
}
