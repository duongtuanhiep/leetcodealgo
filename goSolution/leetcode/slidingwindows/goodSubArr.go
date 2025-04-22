package slidingwindows

/*

Question 2537: https://leetcode.com/problems/count-the-number-of-good-subarrays

Sliding windows: count the possibility via combination counts

**/

func countGood(nums []int, k int) int64 {
	res := 0

	count, currentGood := make(map[int]int), 0
	l, r := 0, 0
	for r < len(nums) {

		// Extend r until it's good, return all possible arrs
		for r < len(nums) && currentGood < k {
			if count[nums[r]] > 0 {
				oldCount := (count[nums[r]] - 1) * count[nums[r]] / 2
				count[nums[r]]++
				currentGood += (count[nums[r]]-1)*count[nums[r]]/2 - oldCount
			} else {
				count[nums[r]]++
			}
			r++
		}

		if currentGood >= k {
			res += len(nums) - r + 1
		}

		// Once it's good, extend l and counts all possible arrs until it's not good
		for l < len(nums) && currentGood >= k {
			if count[nums[l]] > 0 {
				oldCount := (count[nums[l]] - 1) * count[nums[l]] / 2
				count[nums[l]]--
				currentGood += (count[nums[l]]-1)*count[nums[l]]/2 - oldCount
			} else {
				count[nums[l]]--
			}
			if currentGood >= k {
				res += len(nums) - r + 1
			}
			l++
		}
	}

	return int64(res)
}
