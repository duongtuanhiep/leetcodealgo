from typing import List

"""

Question 3931: https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/

"""


class Solution:
    def minOperations(self, nums: List[int]) -> int:
        index = 0
        res = 0
        while index <= len(nums) - 3:
            if not nums[index]:
                res += 1
                nums[index] = 1
                nums[index + 1] = 0 if nums[index + 1] else 1
                nums[index + 2] = 0 if nums[index + 2] else 1
            index += 1

        if not nums[index] or not nums[index + 1] or not nums[index - 1]:
            return -1
        return res


cases = [
    [0, 1, 1, 1, 0, 0],  # 3
    [0, 1, 1, 1],  # -1
    [0, 0, 0, 1, 1 ,0 ,1 ,1 ,0], # 3
]

for c in cases: 
    print(Solution().minOperations(c))
