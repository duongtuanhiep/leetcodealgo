package array

/*
Question 2845: https://leetcode.com/problems/count-of-interesting-subarrays

Magic explanation: https://leetcode.com/problems/count-of-interesting-subarrays/?envType=daily-question&envId=2025-04-25

A better problem 974: https://leetcode.com/problems/subarray-sums-divisible-by-k/editorial/

The key idea for this is to realize that there is a rotating counts/sums as you traverse the array.
Instead of keeping track of individual, you keep track of the COUNT of requirements.

for 2845: keep track of counts of sub arr that has (cnt & modulo)

for 974: keep track of running sum that has certain mod

The first elem cnt[0] of these problems is set to 1 because that's basically reserved for the first subarr (that always satify)
You can consider it a base case
*/
func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	var res int64 = 0
	var prefix int = 0
	cnt[0] = 1
	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			prefix++
		}
		res += int64(cnt[(prefix-k+modulo)%modulo])
		cnt[prefix%modulo]++
	}
	return res
}

func subarraysDivByK(nums []int, k int) int {
	runningSumMap := make(map[int]int)
	res, mod := 0, 0
	runningSumMap[0] = 1
	for _, val := range nums {
		mod = (mod + val%k + k) % k // makes it positive in our arr since (mod -2 == mod 3 for  )
		res += runningSumMap[mod]
		runningSumMap[mod] += 1
	}
	return res
}
