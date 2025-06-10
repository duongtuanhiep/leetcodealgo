class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        dp = [-1e5] * len(nums)
        res = -1e5
        for i, n in nums:
            if i == 0:
                dp[i] = nums[i]
            else:
                dp[i] = nums[i] + max(dp[i - 1], 0)
            res = max(res, dp[i])
        return res
