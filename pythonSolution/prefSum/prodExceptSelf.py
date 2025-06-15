from typing import List

"""
Question 238:
Do 2 pass: 1 forward, 1 backward using rolling fixes
"""


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        # 2 pass on result arrays
        res = [0] * len(nums)
        rollingSum = 1
        for i, n in enumerate(nums):
            rollingSum *= n
            res[i] = rollingSum

        rollingSum = 1
        for i in range(len(nums) - 1, -1, -1):
            if i + 1 < len(nums):
                rollingSum *= nums[i + 1]

            left = 1
            if i - 1 >= 0:
                left = res[i - 1]
            res[i] = left * rollingSum

        return res
