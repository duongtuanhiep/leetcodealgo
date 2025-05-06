from typing import List
from collections import deque

"""

Question 2962: https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times

2 pointers, keep track of current counts of k.
"""


class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        res, i, cnt, maxElem = 0, 0, 0, max(nums)
        for j, val in enumerate(nums):
            if val == maxElem:
                cnt += 1
            while cnt >= k:
                res += len(nums) - j
                if nums[i] == maxElem:
                    cnt -= 1
                i += 1

        return res


if __name__ == "__main__":
    cases = [
        [[1, 3, 2, 3, 3], 2],  # 6
        [[1, 4, 2, 1], 3],  # 0
        [[3, 1, 3, 2, 3, 3], 2],  #
    ]

    for c in cases:
        print(Solution().countSubarrays(c[0], c[1]))
