package slidingwindows

/*

Question 2799: https://leetcode.com/problems/count-complete-subarrays-in-an-array

Typical sliding windows with l,r and res += Possible subarr valid from now (len - right)
**/

func countCompleteSubarrays(nums []int) int {
	allChar := make(map[int]struct{})
	for _, n := range nums {
		allChar[n] = struct{}{}
	}

	l, r, res := 0, 0, 0
	curChar := make(map[int]int)
	for r < len(nums) {
		for ; r < len(nums) && len(curChar) < len(allChar); r++ {
			curChar[nums[r]]++
		}

		if len(curChar) == len(allChar) {
			res += len(nums) - r + 1
		}

		// trimming off l
		for ; l < len(nums) && len(curChar) == len(allChar); l++ {
			curChar[nums[l]]--
			if curChar[nums[l]] == 0 {
				delete(curChar, nums[l])
			}

			if len(curChar) == len(allChar) {
				res += len(nums) - r + 1
			}
		}

	}
	return res
}
