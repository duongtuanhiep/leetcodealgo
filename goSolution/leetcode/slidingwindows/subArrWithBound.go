package slidingwindows

/*
Question 2444: https://leetcode.com/problems/count-subarrays-with-fixed-bounds

Keep track of min bounds when extending and max bound when retracting [l:r]

**/

func countSubarrays(nums []int, minK int, maxK int) int64 {
	res, start, minI, maxI := 0, -1, -1, -1
	for i := range nums {
		if nums[i] < minK || nums[i] > maxK {
			start = i
		}
		if nums[i] == minK {
			minI = i
		}
		if nums[i] == maxK {
			maxI = i
		}
		valid := max(0, min(minI, maxI)-start)
		res += valid
	}
	return int64(res)
}
