import math


class Solution:
    def maxAbsoluteSum(self, nums: List[int]) -> int:
        min, max = 0, 0
        res = 0
        for num in nums:
            min = min + num if min + num < num else num
            max = max + num if max + num > num else num
            curmax = abs(min) if abs(min) > abs(max) else abs(max)
            if curmax > res:
                res = curmax
        return res
