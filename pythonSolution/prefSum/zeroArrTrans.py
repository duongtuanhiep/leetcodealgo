"""

Question 3355: https://leetcode.com/problems/zero-array-transformation-i

Instead of counting and iterate freq one by one (O(N^2)) we precompute number changes via a delta arrays of frequency

"""


class Solution:
    def isZeroArray(self, nums: List[int], queries: List[List[int]]) -> bool:
        interval = [0] * (len(nums) + 1)
        for q in queries:
            interval[q[0]] += 1
            interval[q[1] + 1] -= 1

        cur = 0
        for i in range(len(nums)):
            cur += interval[i]
            if nums[i] > cur:
                return False

        return True
