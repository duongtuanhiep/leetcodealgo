from typing import List

"""

Question 2560: https://leetcode.com/problems/house-robber-iv

Brute force: 
- Generate all possible solutions. 
- Optimize into DP. Still not fast enough 

Use BS: 
"""

# DP
class Solution:
    def minCapability(self, nums: List[int], k: int) -> int:
        steal = []
        for i in range(k + 1):
            steal.append([-1 for _ in range(len(nums))])

        self.steal(nums, steal, 0, k)
        return steal[k][0]

    def steal(self, nums: list[int], dp: list[list[int]], index: int, k: int) -> int:
        if len(nums) - index < (k*2 - 1):
            return -2

        if k == 0:
            return 0

        if index >= len(nums):
            return -2

        if dp[k][index] != -1:
            return dp[k][index]

        now = self.steal(nums, dp, index + 2, k - 1)
        next = self.steal(nums, dp, index + 1, k)

        res = -2
        if now != -2:
            res = max(now, nums[index])
        if next != -2:
            res = min(res, next)

        dp[k][index] = res

        return res

# BS
class Solution:
    def minCapability(self, nums, k):
        # Store the maximum nums value in maxReward.
        min_reward, max_reward = 1, max(nums)
        total_houses = len(nums)

        # Use binary search to find the minimum reward possible.
        while min_reward < max_reward:
            mid_reward = (min_reward + max_reward) // 2
            possible_thefts = 0

            index = 0
            while index < total_houses:
                if nums[index] <= mid_reward:
                    possible_thefts += 1
                    index += 2  # Skip the next house to maintain the non-adjacent condition
                else:
                    index += 1

            if possible_thefts >= k:
                max_reward = mid_reward
            else:
                min_reward = mid_reward + 1

        return min_reward

if __name__ == "__main__":
    cases = [
        [[2, 3, 5, 9], 2],  # 5
        [[2, 7, 9, 3, 1], 2],  # 2
    ]

    for c in cases:
        print(Solution().minCapability(c[0], c[1]))
