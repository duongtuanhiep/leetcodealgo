from typing import List
from collections import deque

"""

Questions 84: https://leetcode.com/problems/largest-rectangle-in-histogram

Constraint:
- len N = 1e5 
- value M = 1e4

Possible solutions:
- For each steps i, finds best j appear before i for max -> O(N^2)
- Move from max heights to 1. Find max at each -> O(N*M)
- Use monotonic stacks to reduce candidate counts at each steps -> O(N * M) but more optimized
"""


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        res = heights[0]
        ms = deque()
        for i, n in enumerate(heights):
            for index, val in ms:
                res = max(res, (i - index + 1) * min(val, n))

            lastIndexWithValueN = i
            while ms and ms[-1][1] >= n:
                lastIndexWithValueN, _ = ms.pop()

            ms.append((lastIndexWithValueN, n))

        return res
