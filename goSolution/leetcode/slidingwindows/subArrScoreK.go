package slidingwindows

/*

Question 2302:https://leetcode.com/problems/count-subarrays-with-score-less-than-k

**/

func countSubarrays(nums []int, k int64) int64 {
	res, total := 0, 0
	for i, j := 0, 0; j < len(nums); j++ {
		total += nums[j]
		for i <= j && total*(j-i+1) >= int(k) {
			total -= nums[i]
			i++
		}
		res += j - i + 1
	}
	return int64(res)
}
