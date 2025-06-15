from typing import List

"""

Question 2616: https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/

Idea: 
- DP doesn't work cause there ain't no rules
- So is "default" greedy
- But we can bound the possible result from range [1:max] and if we can check in O(N), we should be able to do at most ~30 ops to figure out the result
"""


class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        if not p:
            return 0

        hi, lo = abs(max(nums) - min(nums)), 0
        nums.sort()

        def canSatify(nums: list[int], pairCnt: int, target: int) -> bool:
            i = 0
            while i < len(nums) - 1:
                if abs(nums[i] - nums[i + 1]) <= target:
                    pairCnt -= 1
                    i += 2
                else:
                    i += 1
                if not pairCnt:
                    return True
            return False

        res = hi
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if canSatify(nums, p, mid):
                res = mid
                hi = mid
            else:
                lo = mid + 1
        return res
